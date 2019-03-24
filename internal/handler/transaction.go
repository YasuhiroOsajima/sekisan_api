package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"sekisan_api/internal/service"
)

type transactionServiceI interface {
	GetTransactionList() (tList service.TransactionList, err error)
	AddTransaction(employeeNum, hour int, operation, reason string) (res int64, err error)
}

type transactionHandler struct {
	service transactionServiceI
}

func NewTransactionHandler(transactionService transactionServiceI) *transactionHandler {
	return &transactionHandler{
		service: transactionService,
	}
}

// Transaction handler
func (h *transactionHandler) GetTransactionList(w http.ResponseWriter, r *http.Request) {
	transactionList, err := h.service.GetTransactionList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(transactionList); err != nil {
		panic(err)
	}
}

func (h *transactionHandler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	hourStr := v["hour"]
	operation := v["operation"]
	reason := v["reason"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	a, err := h.service.AddTransaction(employeeNum, hour, operation, reason)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}
