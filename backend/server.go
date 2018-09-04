package backend

import (
	"fmt"
	"net/http"
)
var httpListenAndServe = http.ListenAndServe

func RunServer() (err error) {
	fs := http.FileServer(http.Dir("./frontend/dist/gospa"))
	http.Handle("/", fs)
	fmt.Println("Listening on http://127.0.0.1:8080")
	err = httpListenAndServe(":8080", nil)
	return
}
