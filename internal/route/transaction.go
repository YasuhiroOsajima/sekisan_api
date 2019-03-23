package route

import (
	"github.com/gorilla/mux"
	"sekisan_api/internal/handler"
	"sekisan_api/internal/repository"
	"sekisan_api/internal/service"
)

func AddTransactionRoute(r *mux.Router) {
	tr := repository.NewTransactionRepository()
	ts := service.NewTransactionService(tr)
	t := handler.NewTransactionHandler(ts)

	r.HandleFunc("/sekisan", t.GetTransactionList).Methods("GET")
	r.HandleFunc("/sekisan", t.AddTransaction).Methods("POST")
}
