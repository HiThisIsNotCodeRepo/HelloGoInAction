package main

import (
	"github.com/qinchenfeng/HelloGoInAction/Chap9/endpoint/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
