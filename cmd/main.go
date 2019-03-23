package cmd

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"sekisan_api/configs"
	"sekisan_api/internal/handler"
	"sekisan_api/internal/route"
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

	// Add handlers.
	route.AddAdminRoute(r)
	route.AddMemberRoute(r)
	route.AddSekisanRoute(r)
	route.AddTransactionRoute(r)

	// Now found handler.
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	// Start HTTP server.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://127.0.0.1:5000"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
	})
	hdr := c.Handler(r)
	addr := ":" + configs.Port

	log.Printf("[INFO] start server %s", addr)
	log.Fatal(http.ListenAndServe(addr, context.ClearHandler(
		handlers.LoggingHandler(os.Stderr, hdr))))
}
