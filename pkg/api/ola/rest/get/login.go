package get

import (
	controller "github.com/ola/pkg/controllers"
	"github.com/ola/pkg/handler"
	"github.com/valyala/fasthttp"
)

//Login function to login user
func Login(ctx *fasthttp.RequestCtx) {

	params := []string{"username", "pass"}
	parameters := controller.GetParameters(ctx, params)
	user := handler.GetUser(parameters["username"])

	if user != nil {
		if user.Password == parameters["pass"] {
			handler.UpdateLoggedStatus(user.Username, true)
			controller.SucessResponse(ctx, "pass")
		} else {
			controller.ErrorResponse(ctx, "Given password fot this user is wrong, please try again.")
		}
	} else {
		controller.ErrorResponse(ctx, "This user do not exist")
	}
}
