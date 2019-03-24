package service

import (
	"log"

	"sekisan_api/internal/repository"
)

type TransactionList struct {
	Transaction []repository.Transaction
}

type transactionRepositoryI interface {
	GetTransaction() (tl []repository.Transaction, err error)
	AddTransaction(employeeNum, hour int, operation, reason string) (int64, error)
}

type transactionService struct {
	tRepository transactionRepositoryI
}

func NewTransactionService(tr transactionRepositoryI) *transactionService {
	return &transactionService{
		tRepository: tr,
	}
}

func (s *transactionService) GetTransactionList() (tList TransactionList, err error) {
	tl, err := s.tRepository.GetTransaction()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	tList = TransactionList{tl}
	return
}

func (s *transactionService) AddTransaction(employeeNum, hour int, operation, reason string) (res int64, err error) {
	res, err = s.tRepository.AddTransaction(employeeNum, hour, operation, reason)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	return
}
