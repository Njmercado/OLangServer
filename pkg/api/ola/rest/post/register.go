package post

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

//SignUp : function to register that given user
func SignUp(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "user has been signed up")
}