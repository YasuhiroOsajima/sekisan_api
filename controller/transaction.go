package controller

import (
	"log"
	"sekisan_api/model"
)

type transactionList struct {
	Transaction []model.Transaction
}

func getTransactionList() (tList transactionList, err error) {
	tl, err := model.GetTransaction()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	tList = transactionList{tl}
	return
}

func addTransaction(employeeNum, hour int, operation, reason string) (t model.Transaction, err error) {
	t, err = model.AddTransaction(employeeNum, hour, operation, reason)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	return
}
