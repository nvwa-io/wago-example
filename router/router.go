package router

import (
	"github.com/nvwa-io/wago"
	"github.com/nvwa-io/wago-example/controller"
	"github.com/nvwa-io/wago-example/controller/home"
	"github.com/nvwa-io/wago/middleware"
	"log"
)

func init() {
	log.Println("Router inited")

	wago.Use(middleware.RequestId(),
		middleware.RequestLogger())

	group := wago.NewRouterGroup().
		Prefix("/v1").
		Controller(
			new(controller.ExampleController),
			new(home.Example2Controller))

	wago.AddRouterGroups(group)
}
