package controller

import "net/http"

func notFound(w http.ResponseWriter) {
	code := http.StatusNotFound
	http.Error(w, http.StatusText(code), code)
}

func badRequest(w http.ResponseWriter) {
	code := http.StatusBadRequest
	http.Error(w, http.StatusText(code), code)
}

func forbidden(w http.ResponseWriter) {
	code := http.StatusForbidden
	http.Error(w, http.StatusText(code), code)
}
