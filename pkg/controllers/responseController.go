package controller

import "github.com/valyala/fasthttp"

func SucessResponse(ctx *fasthttp.RequestCtx, message string) {

	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	response := []byte(`{
		"response":` + message + ` 
	}`)
	ctx.Response.SetBody(response)
}

func ErrorResponse(ctx *fasthttp.RequestCtx, message string) {

	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(400)
	response := []byte(`{
		"response":` + message + ` 
	}`)
	ctx.Response.SetBody(response)
}
