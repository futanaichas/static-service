package main

import (
	"net/http"

	"github.com/futanaichas/static-service/src/handle"
)

func main() {
	http.Handle("/static/", handle.Taken())
	http.Handle("/api/static/upload", handle.Upload())
	http.ListenAndServe(":80", nil)
}
