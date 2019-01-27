package controller

import (
	"log"
	"sekisan_api/model"
)

type transactionList struct {
	Transaction []model.Transaction
}

func getTransactionList() (memberList, error) {
	tList, err := model.GetTransaction()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	memberList := transactionList{tList}
	return memberList, err
}

