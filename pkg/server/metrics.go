package server

import "net/http"

func http_handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
