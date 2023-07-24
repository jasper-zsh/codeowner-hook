package features

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jasper-zsh/codeowner-hook/app"
	"github.com/jasper-zsh/codeowner-hook/app/providers/github"
	"github.com/jasper-zsh/codeowner-hook/app/utils"
	"github.com/sirupsen/logrus"
)

type CodeOwnerRule struct {
	Rule   string
	Owners utils.Set[string]

	patternLevels []string
}

func NewCodeOwnerRuleFromLine(line string) *CodeOwnerRule {
	line = strings.Trim(line, " ")
	if strings.HasPrefix(line, "#") {
		return nil
	}
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return nil
	}
	rule := &CodeOwnerRule{
		Rule: line,
	}

	pattern := parts[0]
	rule.Owners = utils.Set[string]{}
	rule.Owners.Add(parts[1:]...)
	rule.patternLevels = strings.Split(pattern, "/")
	if rule.patternLevels[0] == "" {
		rule.patternLevels = rule.patternLevels[1:]
	}
	return rule
}

func (r *CodeOwnerRule) Match(file string) bool {
	fileLevels := strings.Split(file, "/")
	if len(fileLevels) < len(r.patternLevels) {
		return false
	}
	filePos, patternPos := 0, 0
	for filePos < len(fileLevels) && patternPos < len(r.patternLevels) {
		if r.patternLevels[patternPos] == "**" {
			patternPos++
			if patternPos == len(r.patternLevels) {
				break
			}
			for ; filePos < len(fileLevels) && fileLevels[filePos] != r.patternLevels[patternPos]; filePos++ {
			}
			if filePos == len(fileLevels) {
				break
			}
		} else if r.patternLevels[patternPos] != "*" && r.patternLevels[patternPos] != fileLevels[filePos] {
			return false
		}
		filePos++
		patternPos++
	}
	if patternPos == len(r.patternLevels) {
		return true
	} else {
		return false
	}
}

func CheckCodeOwner(event github.GithubPushEvent) {
	rawCodeOwners, err := getCodeOwnersFile(event.Repository.Owner.Name, event.Repository.Name)
	if err != nil {
		logrus.Errorf("Failed to get CODEOWNERS %+v", err)
		return
	}

	authorChanges := make(map[string]utils.Set[string])
	matchedChangeAuthorsMap := make(map[string]utils.Set[string])
	affectedOwnerSet := utils.Set[string]{}
	for _, commit := range event.Commits {
		changes, ok := authorChanges[commit.Author.Email]
		if !ok {
			changes = utils.Set[string]{}
			authorChanges[commit.Author.Email] = changes
		}
		changes.Add(commit.Added...)
		changes.Add(commit.Modified...)
		changes.Add(commit.Removed...)
	}

	rows := strings.Split(rawCodeOwners, "\n")
	for _, row := range rows {
		rule := NewCodeOwnerRuleFromLine(row)
		if rule == nil {
			continue
		}

		for author, changes := range authorChanges {
			for change := range changes {
				if rule.Match(change) && !rule.Owners.Contains(author) {
					matchedChanges, ok := matchedChangeAuthorsMap[change]
					if !ok {
						matchedChanges = utils.Set[string]{}
						matchedChangeAuthorsMap[change] = matchedChanges
					}
					matchedChanges.Add(author)
					affectedOwnerSet.Add(rule.Owners.ToArray()...)
				}
			}
		}
	}
	logrus.Infof("matchedChanges %+v", matchedChangeAuthorsMap)
	if len(affectedOwnerSet) > 0 && len(matchedChangeAuthorsMap) > 0 {
		err = sendQyWxBot(event.Repository.FullName, event.Ref, affectedOwnerSet.ToArray(), matchedChangeAuthorsMap)
		if err != nil {
			logrus.Errorf("Failed to send qyweixin message %+v", err)
		}
	}
}

func sendQyWxBot(repo, branch string, owners []string, changes map[string]utils.Set[string]) error {
	ownersTxt := make([]string, 0, len(owners))
	for _, owner := range owners {
		ownersTxt = append(ownersTxt, fmt.Sprintf("> Owner: %s", owner))
	}
	changesTxt := make([]string, 0, len(changes))
	for change, authors := range changes {
		changesTxt = append(changesTxt, fmt.Sprintf("%s 修改者: %s", change, strings.Join(authors.ToArray(), ", ")))
	}
	payload := map[string]any{
		"msgtype": "markdown",
		"markdown": map[string]any{
			"content": fmt.Sprintf("%d位Owner的代码发生%d个变更，请关注！\n> 仓库: %s\n> 分支: %s\n%s\n%s",
				len(owners),
				len(changes),
				repo,
				branch,
				strings.Join(ownersTxt, "\n"),
				strings.Join(changesTxt, "\n"),
			),
		},
	}
	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Post(app.Config.QyWeixinBot, "application/json", bytes.NewReader(rawPayload))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("http error %d: %s", resp.StatusCode, body)
	}
	logrus.Infof("qyweixin bot response: %s", body)
	return nil
}

func getCodeOwnersFile(owner string, repo string) (content string, err error) {
	paths := []string{"docs/CODEOWNERS", ".github/CODEOWNERS"}
	for _, p := range paths {
		fileContent, _, resp, err := github.GithubClient.Repositories.GetContents(context.Background(), owner, repo, p, nil)
		if err != nil {
			if resp.StatusCode == 404 {
				continue
			}
			logrus.Errorf("failed to get CODEOWNERS %+v", err)
			return "", err
		}
		if fileContent != nil {
			content, err = fileContent.GetContent()
			if err != nil {
				return "", err
			}
		}
	}
	return
}
