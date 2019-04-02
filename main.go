package main

import (
	"github.com/nvwa-io/wago"
	_ "github.com/nvwa-io/wago-example/router"
)

func main() {
	wago.Serve()
}
