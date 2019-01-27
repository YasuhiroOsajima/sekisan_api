package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"sekisan_api/config"
	"sekisan_api/controller"
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Panicln(err)
	}
	time.Local = loc
}

func main() {
	r := mux.NewRouter()
	h := controller.NewHandler()

	// Admin handlers.
	r.HandleFunc("/sekisan_app/admin", h.GetAdminList).Methods("GET")
	r.HandleFunc("/sekisan_app/admin", h.RegisterAdmin).Methods("POST")
	r.HandleFunc("/sekisan_app/{admin_id[0-9]+}/name", h.UpdateAdminName).Methods("POST")
	r.HandleFunc("/sekisan_app/{admin_id[0-9]+}/name", h.UpdateAdminPassword).Methods("POST")
	r.HandleFunc("/sekisan_app/{admin_id[0-9]+}/name", h.UpdateAdminEnabled).Methods("POST")

	// Member handlers.
	r.HandleFunc("/sekisan_app/member", h.GetMemberList).Methods("GET")
	r.HandleFunc("/sekisan_app/member", h.RegisterMember).Methods("POST")
	r.HandleFunc("/sekisan_app/member", h.UpdateMemberName).Methods("POST")
	r.HandleFunc("/sekisan_app/member", h.UpdateMemberEnabled).Methods("POST")

	// Sekisan handlers.
	r.HandleFunc("/sekisan", h.GetAllSekisan).Methods("GET")

	// Transaction handlers.

	// Now found handler.
	r.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)

	// Start HTTP server.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://127.0.0.1:5000"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
	})
	hdr := c.Handler(r)
	addr := ":" + config.Port

	log.Printf("[INFO] start server %s", addr)
	log.Fatal(http.ListenAndServe(addr, context.ClearHandler(
		handlers.LoggingHandler(os.Stderr, hdr))))
}
