package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/v53/github"
	"github.com/jasper-zsh/codeowner-hook/app"
	"golang.org/x/oauth2"
)

var (
	GithubClient *github.Client
)

func InitGithubClient() {
	var cli *http.Client
	if app.Config.GithubToken != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: app.Config.GithubToken})
		cli = oauth2.NewClient(context.Background(), ts)
	}
	GithubClient = github.NewClient(cli)
}
