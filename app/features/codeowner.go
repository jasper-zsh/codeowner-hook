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

func CheckCodeOwner(event github.GithubPushEvent) {
	rawCodeOwners, err := getCodeOwnersFile(event.Repository.Owner.Name, event.Repository.Name)
	if err != nil {
		logrus.Errorf("Failed to get CODEOWNERS %+v", err)
		return
	}

	changes := utils.Set[string]{}
	matchedChanges := make(map[string]utils.Set[string])
	for _, commit := range event.Commits {
		changes.Add(commit.Added...)
		changes.Add(commit.Modified...)
		changes.Add(commit.Removed...)
	}

	rows := strings.Split(rawCodeOwners, "\n")
	for _, row := range rows {
		row = strings.Trim(row, " ")
		if strings.HasPrefix(row, "#") {
			continue
		}
		parts := strings.Split(row, " ")
		if len(parts) < 2 {
			continue
		}
		pattern := parts[0]
		owners := utils.Set[string]{}
		owners.Add(parts[1:]...)
		patternLevels := strings.Split(pattern, "/")
		if patternLevels[0] == "" {
			patternLevels = patternLevels[1:]
		}
		logrus.Infof("pathLevels %+v", patternLevels)

		for change := range changes {
			if matchPatternLevels(change, patternLevels) {
				for owner := range owners {
					m, ok := matchedChanges[owner]
					if !ok {
						m = utils.Set[string]{}
						matchedChanges[owner] = m
					}
					m.Add(change)
				}
			}
		}
	}
	logrus.Infof("matchedChanges %+v", matchedChanges)
	owners := make([]string, 0, len(matchedChanges))
	matchedChangesSet := utils.Set[string]{}
	for owner, matched := range matchedChanges {
		owners = append(owners, owner)
		for change := range matched {
			matchedChangesSet.Add(change)
		}
	}
	err = sendQyWxBot(event.Repository.FullName, event.Ref, owners, matchedChangesSet.ToArray())
	if err != nil {
		logrus.Errorf("Failed to send qyweixin message %+v", err)
	}
}

func sendQyWxBot(repo, branch string, owners, changes []string) error {
	ownersTxt := make([]string, 0, len(owners))
	for _, owner := range owners {
		ownersTxt = append(ownersTxt, fmt.Sprintf("> Owner: %s", owner))
	}
	changesTxt := make([]string, 0, len(changes))
	for _, change := range changes {
		changesTxt = append(changesTxt, fmt.Sprintf("变更: %s", change))
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

func matchPatternLevels(file string, patternLevels []string) bool {
	fileLevels := strings.Split(file, "/")
	if len(fileLevels) < len(patternLevels) {
		return false
	}
	filePos, patternPos := 0, 0
	for filePos < len(fileLevels) && patternPos < len(patternLevels) {
		if patternLevels[patternPos] == "**" {
			patternPos++
			if patternPos == len(patternLevels) {
				break
			}
			for ; filePos < len(fileLevels) && fileLevels[filePos] != patternLevels[patternPos]; filePos++ {
			}
			if filePos == len(fileLevels) {
				break
			}
		} else if patternLevels[patternPos] != "*" && patternLevels[patternPos] != fileLevels[filePos] {
			return false
		}
		filePos++
		patternPos++
	}
	if patternPos == len(patternLevels) {
		return true
	} else {
		return false
	}
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
		}
	}
	return
}
