package main

import (
	"github.com/jasper-zsh/codeowner-hook/app"
	githubProvider "github.com/jasper-zsh/codeowner-hook/app/providers/github"
	"github.com/jasper-zsh/codeowner-hook/app/webhooks/github"
	"github.com/kataras/iris/v12"
)

func main() {
	app.InitConfig()
	githubProvider.InitGithubClient()

	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World!")
	})

	webhooksAPI := app.Party("/webhooks")
	webhooksAPI.Post("/github", github.GithubWebhook)

	app.Listen(":4567")
}
