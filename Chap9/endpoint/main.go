// 这个示例程序实现了简单的网络服务
package main

import (
	"github.com/qinchenfeng/HelloGoInAction/Chap9/endpoint/handlers"
	"log"
	"net/http"
)

// main 是应用程序的入口
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
