package home

import (
	"github.com/nvwa-io/wago"
	"log"
)

// supported router syntax
// @/v1 , All kind of HTTP request method are allowed.
// @get /url/path , "/url/path" allows HTTP GET METHOD.
// @get,post /url/path , "/url/path" allows HTTP GET or POST METHOD.

// @prefix /example2
type Example2Controller struct {
	wago.Controller
}

// @TODO hello world
// @get,post,put /hello
func (t *Example2Controller) HelloWorld() {
	log.Println(t.Ctx.Request.Method)
	log.Println("Hello world")
}

// @TODO Hello world 2
// @/hello/world
func (t *Example2Controller) HelloWorld2() {
	log.Println("Hello world2")
}

// @Get,Post /hello-world
func (t Example2Controller) HelloWorld3() {
	log.Println("Hello world3")
}

// @Post /hello-world4
func (t Example2Controller) HelloWorld4() {
	log.Println("Hello world4")
}
