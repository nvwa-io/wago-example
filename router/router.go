package router

import (
	"github.com/nvwa-io/wago"
	"github.com/nvwa-io/wago-example/controller"
	"github.com/nvwa-io/wago-example/controller/home"
	"log"
)

func init() {
	log.Println("Router inited")

	group := wago.NewRouterGroup().
		Prefix("/v1").
		//Use().
		Controller(
			new(controller.ExampleController),
			new(home.Example2Controller))

	wago.AddRouterGroups(group)
}
