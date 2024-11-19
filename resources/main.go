package resources

import (
	"github.com/nitrictech/go-sdk/nitric"
	"github.com/nitrictech/go-sdk/nitric/apis"
	"github.com/nitrictech/go-sdk/nitric/topics"
)

type Resource struct {
	MainApi     apis.Api
	NotifyTopic topics.SubscribableTopic
}

var resource = &Resource{
	MainApi:     nitric.NewApi("test"),
	NotifyTopic: nitric.NewTopic("user-created"),
}

func Get() *Resource {
	return resource
}
