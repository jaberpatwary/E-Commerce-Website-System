package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type TransactionreportRepositoryInterface interface {
	InsertTransactionreport(model.PostTransactionreport) bool
	GetAllTransactionreport() []model.Transactionreport
	GetOneTransactionreport(int) model.Transactionreport
	DeleteTransactionreport(int) bool
	UpdateTransactionreport(int, model.PostTransactionreport) model.Transactionreport
}
