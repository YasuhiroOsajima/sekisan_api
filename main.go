package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"sekisan_api/config"
	"sekisan_api/controller"
	"time"

	"github.com/rs/cors"
	"sekisan_api/model"
)

const (
	SessionSecret = "testtest"
)

func init() {
	var err error
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Panicln(err)
	}
	time.Local = loc
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Found\n"))
}

func main() {
	model.Db_connect()
	store := sessions.NewCookieStore([]byte(SessionSecret))

	r := mux.NewRouter()
	h := controller.NewHandler(store)

	// Add handlers.
	r.HandleFunc("/sekisan/{emp_id:[0-9]+}", h.GetSekisan).Methods("GET")
	r.HandleFunc("/sekisan", h.GetAllSekisan).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// Start HTTP server.
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:8080", "http://127.0.0.1:5000"},
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		Debug: true,
	})
	hdr := c.Handler(r)
	addr := ":" + config.Port

	log.Printf("[INFO] start server %s", addr)
	log.Fatal(http.ListenAndServe(addr, context.ClearHandler(
		handlers.LoggingHandler(os.Stderr, hdr))))
}
