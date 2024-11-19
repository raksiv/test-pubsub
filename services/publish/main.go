package main

import (
	"context"
	"fmt"

	"github.com/nitrictech/go-sdk/nitric"
	"github.com/nitrictech/go-sdk/nitric/apis"
	"github.com/nitrictech/go-sdk/nitric/topics"
	"github.com/nitrictech/templates/go-starter/resources"
)

func main() {

	var publishTopic = resources.Get().NotifyTopic.Allow(topics.TopicPublish)

	resources.Get().MainApi.Get("/test", func(ctx *apis.Ctx) {

		err := publishTopic.Publish(context.TODO(), map[string]interface{}{
			"email": "new.user@example.com",
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("triggered ... ")
	})

	nitric.Run()
}
