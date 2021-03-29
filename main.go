package main

import (
	"github.com/kataras/iris/v12"
	"iris-test/routes"
)

func main() {
	app := iris.Default()
	routes.Routes(app)
	//fmt.Printf("%T", app)
	_ = app.Listen(":4700")
}
