package interceptor

import (
	"github.com/jinzhu/gorm"
	"github.com/lios/go-activiti/db"
)

type TransactionContextInterceptor struct {
	Next CommandInterceptor
}

func (transactionContextInterceptor TransactionContextInterceptor) Execute(command Command) (value interface{}, err error) {
	defer db.ClearTXDB()
	db.GORM_DB.Transaction(func(tx *gorm.DB) error {
		db.InitTXDB(tx)
		value, err = transactionContextInterceptor.Next.Execute(command)
		return err
	})
	return value, err
}

func (transactionContextInterceptor *TransactionContextInterceptor) SetNext(next CommandInterceptor) {
	transactionContextInterceptor.Next = next
}
