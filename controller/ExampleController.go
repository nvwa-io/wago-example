package controller

import (
	"github.com/nvwa-io/wago"
	"log"
)

// supported router syntax
// @/v1 , All kind of HTTP request method are allowed.
// @get /url/path , "/url/path" allows HTTP GET METHOD.
// @get,post /url/path , "/url/path" allows HTTP GET or POST METHOD.

// @prefix /example
type ExampleController struct {
	wago.Controller
}

// Get, Post or Put supported
// @get,post,put /hello
func (t *ExampleController) HelloWorld() {
	log.Println(t.Ctx.Request.Method)
	log.Println("Hello world")
}

// Get or Post
// @/hello/world
func (t *ExampleController) HelloWorld2() {
	log.Println("Hello world2")
}

func (t *ExampleController) Default() {
	log.Println("-> in default")
}

func (t *ExampleController) Hello_POST() {
	log.Println("-> in Hello_POST")
}

func (t *ExampleController) GetTaskByGid() {

}

func (t *ExampleController) Hello_Get() {
	log.Println("-> in Hello_Get().")
}

func (t *ExampleController) Hello_PUT() {
	log.Println("-> in Hello_PUT().")
}

// @Get,Post /hello-world
func (t ExampleController) HelloWorld3() {
	log.Println("Hello world3")
}

// @Post /hello-world4
func (t ExampleController) HelloWorld4() {
	log.Println("Hello world4")
}
