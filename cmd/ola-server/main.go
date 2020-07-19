package main

import (
	"github.com/ola/pkg/api/ola/rest"
)

func main() {
	rest.NewServer(":8000")
}
