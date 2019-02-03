package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	sessionSecret = "testtest"
	sessionName   = "test_session"
)

//var BaseTime time.Time
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	store = sessions.NewCookieStore([]byte(sessionSecret))
)

// Correct response handlers.
type Handler struct {
	store sessions.Store
}

func NewHandler() *Handler {
	//BaseTime = time.Date(2018, 10, 16, 10, 0, 0, 0, time.Local)
	return &Handler{
		store: store,
	}
}

// 404 handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Not Found\n"))
}

// Response utility
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
