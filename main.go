package main

import (
	"github.com/nvwa-io/wago"
	_ "github.com/nvwa-io/wago-example/router"
)

func main() {
	//wago.WagoApp.Server.Handle("GET", "/test",
	//	func(context *gin.Context) {
	//		log.Println("111111")
	//		context.AbortWithStatusJSON(500, map[string]interface{}{
	//			"hello": "world",
	//		})
	//	},
	//
	//	middleware.RequestId(),
	//	func(context *gin.Context) {
	//		log.Println("22222")
	//	})
	//wago.WagoApp.Server.Use(wago.MiddlewareRequestId())
	//wago.Use(middleware.RequestId())

	wago.Serve()
}
