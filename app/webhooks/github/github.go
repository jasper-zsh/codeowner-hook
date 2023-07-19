package github

import (
	"github.com/jasper-zsh/codeowner-hook/app/features"
	githubProvider "github.com/jasper-zsh/codeowner-hook/app/providers/github"
	"github.com/jasper-zsh/codeowner-hook/app/utils"
	"github.com/kataras/iris/v12"
)

func GithubWebhook(ctx iris.Context) {
	action := ctx.GetHeader("X-GitHub-Event")
	rawEvent := make(map[string]any)
	err := ctx.ReadJSON(&rawEvent)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	event := githubProvider.GithubEvent{}
	err = utils.DecodeMap(rawEvent, &event)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	switch action {
	case "push":
		pushEvent := githubProvider.GithubPushEvent{
			GithubEvent: event,
		}
		err = utils.DecodeMap(event.Other, &pushEvent)
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}
		go features.CheckCodeOwner(pushEvent)
	}
	ctx.StatusCode(iris.StatusOK)
}
