package main

import (
	"fmt"

	"github.com/nitrictech/go-sdk/nitric"
	"github.com/nitrictech/go-sdk/nitric/topics"
	"github.com/nitrictech/templates/go-starter/resources"
)

func main() {

	resources.Get().NotifyTopic.Subscribe(func(ctx *topics.Ctx) {
		email := ctx.Request.Message()["email"].(string)

		fmt.Println("The email is: " + email)
	})

	nitric.Run()
}
