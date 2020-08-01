package post

import (
	"errors"

	"github.com/ola/pkg/controller"
	"github.com/ola/pkg/handler"
	"github.com/valyala/fasthttp"
)

func areValidUsernameAndPassword(username, password string) error {

	if !controller.IsValidUsername(username) {
		return errors.New("Invalid username. Remember than a valid username must have numbers, lowercase and uppercase letter and minimun 6 characters")
	}

	if !controller.IsValidPassword(password) {
		return errors.New(`Invalid password. Remember than a valid password must have numbers, lowecase and uppercase letter. 
		Also, it must have special characters and a minimun lenght of 8`)
	}

	return nil
}

//SignUp : function to register that given user
func SignUp(ctx *fasthttp.RequestCtx) {
	params := []string{"username", "pass"}
	parameters := controller.GetParameters(ctx, params)
	user, _ := handler.GetUser(parameters["username"])

	if user != nil {
		controller.ErrorResponse(ctx, "This user already exist")
		return
	}

	err := areValidUsernameAndPassword(parameters["username"], parameters["pass"])
	if err != nil {
		controller.ErrorResponse(ctx, err.Error())
		return
	}

	handler.CreateUser(parameters["username"], parameters["pass"])
	controller.SucessResponse(ctx, "User has been inserted correctly")
	return
}
