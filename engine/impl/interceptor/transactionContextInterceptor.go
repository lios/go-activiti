package interceptor

import (
	"github.com/jinzhu/gorm"
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/logger"
)

type TransactionContextInterceptor struct {
	Next CommandInterceptor
}

func (transactionContextInterceptor TransactionContextInterceptor) Execute(command Command) (value interface{}, err error) {
	defer db.ClearTXDB()
	db.GORM_DB.Transaction(func(tx *gorm.DB) error {
		db.InitTXDB(tx)
		value, err = transactionContextInterceptor.Next.Execute(command)
		logger.Error("err:", err)
		return err
	})
	return value, err
}

func (transactionContextInterceptor *TransactionContextInterceptor) SetNext(next CommandInterceptor) {
	transactionContextInterceptor.Next = next
}
