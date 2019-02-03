package route

import (
	"github.com/gorilla/mux"
	"sekisan_api/handler"
	"sekisan_api/repository"
	"sekisan_api/service"
)

func AddTransactionRoute(r *mux.Router) {
	tr := repository.NewTransactionRepository()
	ts := service.NewTransactionService(tr)
	t := handler.NewTransactionHandler(ts)

	r.HandleFunc("/sekisan", t.GetTransactionList).Methods("GET")
	r.HandleFunc("/sekisan", t.AddTransaction).Methods("POST")
}
