package main

import "github.com/kataras/iris/v12"
import "fmt"

func main() {
	app := iris.New()
	fmt.Printf("%T", app)
}