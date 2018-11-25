package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"sekisan_api/config"
	"sekisan_api/controller"
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
	w.Write([]byte("Gorilla!\nNot Found\n"))
}

func status(code int, allow ...string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(code)
		if len(allow) > 0 {
			w.Write([]byte(`Allow: ` + strings.Join(allow, ", ") + "\n"))
		}
	}
}

func main() {
	db := model.Db_connect()
	store := sessions.NewCookieStore([]byte(SessionSecret))

	r := mux.NewRouter()
	h := controller.NewHandler(db, store)

	// Add handlers.
	r.HandleFunc("/sekisan/{emp_id:[0-9]+}", h.GetSekisan).Methods("GET")
	r.HandleFunc("/sekisan/{emp_id:[0-9]+}", status(405, "GET"))
	//r.HandleFunc("/user/{id:[0-9]+}", status(405, "GET")).Methods("POST","PUT", "PATCH", "DELETE")

	r.HandleFunc("/sekisan", h.GetAllSekisan).Methods("GET")
	r.HandleFunc("/sekisan", status(405, "GET"))

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// Start HTTP server.
	addr := ":" + config.Port
	log.Printf("[INFO] start server %s", addr)
	log.Fatal(http.ListenAndServe(addr, context.ClearHandler(handlers.LoggingHandler(os.Stderr, r))))
}
