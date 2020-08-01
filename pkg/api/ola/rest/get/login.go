package get

import (
	"github.com/ola/pkg/controller"
	"github.com/ola/pkg/handler"
	"github.com/valyala/fasthttp"
)

//Login function to login user
func Login(ctx *fasthttp.RequestCtx) {

	params := []string{"username", "pass"}
	parameters := controller.GetParameters(ctx, params)
	user, err := handler.GetUser(parameters["username"])

	if user != nil {
		if user.Password == parameters["pass"] {
			err = handler.UpdateLoggedStatus(user.Username, true)
			controller.SucessResponse(ctx, "pass")
		} else {
			controller.ErrorResponse(ctx, "Given password for this user is wrong, please try again.")
		}
	} else {
		controller.ErrorResponse(ctx, err.Error())
	}
}
