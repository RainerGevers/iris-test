package pongs

import "github.com/kataras/iris/v12"

func GetPong(ctx iris.Context) {
	response := iris.Map{"ping": "pong"}
	_, _ = ctx.JSON(response)
}
