package services

import (
	"github.com/anuragjha/bhayia/services/hello"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/hello", hello.HelloHandler)
}
