package pongs

import "github.com/kataras/iris/v12"

func GetPong(ctx iris.Context) {
	response := iris.Map{"ping": "pong"}
	ctx.StatusCode(214)
	_, _ = ctx.JSON(response)
}
