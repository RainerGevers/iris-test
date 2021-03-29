package routes

import (
	"github.com/kataras/iris/v12"
	"iris-test/app/controllers/v1_0/pongs"
)

func v10(app *iris.Application) {
	// Unauthenticated
	v10 := app.Party("/v1.0")
	{
		v10.Get("/ping", pongs.GetPong)
	}
}
