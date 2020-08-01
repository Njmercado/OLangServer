package get

import (
	"github.com/ola/pkg/config/env"
	"github.com/ola/pkg/controller"
	"github.com/ola/pkg/handler"
	"github.com/valyala/fasthttp"
)

func getDecryptedPassword(password string) string {
	return string(controller.Decrypt([]byte(password), env.GetPassPhrase()))
}

//Login function to login user
func Login(ctx *fasthttp.RequestCtx) {

	params := []string{"username", "pass"}
	parameters := controller.GetParameters(ctx, params)
	user, err := handler.GetUser(parameters["username"])

	if user != nil {

		password := getDecryptedPassword(user.Password)

		if password == parameters["pass"] {
			err = handler.UpdateLoggedStatus(user.Username, true)
			if err != nil {
				controller.ErrorResponse(ctx, err.Error())
			} else {
				controller.SucessResponse(ctx, "pass")
			}
		} else {
			controller.ErrorResponse(ctx, "Given password for this user is wrong, please try again.")
		}
	} else {
		controller.ErrorResponse(ctx, err.Error())
	}
}
