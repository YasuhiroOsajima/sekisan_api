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

	//dbusrpass := dbuser
	//if dbpass != "" {
	//	dbusrpass += ":" + dbpass
	//}

	//dsn := fmt.Sprintf(`%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4`, dbusrpass, dbhost, dbport, dbname)
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	log.Fatalf("mysql connect failed. err: %s", err)
	//}

	store := sessions.NewCookieStore([]byte(SessionSecret))
	//h := controller.NewHandler(db, store)

	h := controller.NewHandler(store)
	r := mux.NewRouter()
	r.HandleFunc("/user/{id:[0-9]+}", h.ShowId).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", status(405, "GET"))
	//r.HandleFunc("/user/{id:[0-9]+}", status(405, "GET")).Methods("POST","PUT", "PATCH", "DELETE")
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	addr := ":" + config.Port
	log.Printf("[INFO] start server %s", addr)
	log.Fatal(http.ListenAndServe(addr, context.ClearHandler(handlers.LoggingHandler(os.Stderr, r))))
}
