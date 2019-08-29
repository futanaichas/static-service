package handle

import "net/http"

// Taken ...
func Taken() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static")))
}
