package main

import (
	"log"
	"net/http"
	"io"
	"os"
	"strings"
	"m2homework/parser"
)

// go mod init xxx xxx是项目的名称。
// 之后需要修改 import 的方式，替换成 import xxx/learn1 的形式。
func main() {
	http.HandleFunc("/header", parser.HeaderHandler(w http.ResponseWriter, r *http.Request))
	http.HandleFunc("/healthz", parser.Healthz())
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
