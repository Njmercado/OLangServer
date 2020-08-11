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

func areEqual(item1, item2 string) bool {
	return item1 == item2
}

//Login function to login user
func Login(ctx *fasthttp.RequestCtx) {

	// Params definition
	params := []string{"username", "pass"}

	// Split those params from url
	parameters := controller.GetParameters(ctx, params)

	//Search for those params into database
	user, err := handler.GetUser(parameters["username"])

	if user != nil {

		//Decrypt password from database
		password := getDecryptedPassword(user.Password)

		if areEqual(password, parameters["pass"]) {

			if err = handler.UpdateLoggedStatus(user.Username, true); err != nil {
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
