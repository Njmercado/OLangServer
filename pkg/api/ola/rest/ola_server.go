package rest

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/ola/pkg/api/ola/rest/get"
	"github.com/ola/pkg/api/ola/rest/post"
	"github.com/valyala/fasthttp"
)

//NewServer this function work to get a new instance of this router
func NewServer(port string) {

	router := fasthttprouter.New()
	router.GET("/login/:username/:pass", get.Login)
	router.POST("/signup", post.SignUp)

	log.Print("server running on port " + port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
