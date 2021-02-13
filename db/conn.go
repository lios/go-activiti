package db

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lios/go-activiti/common"
	"github.com/lios/go-activiti/logger"
	"github.com/lios/go-activiti/runtime"
	"sync"
)

var GORM_DB *gorm.DB

var TXDB *sync.Map

var DNS string

var ROOT string

const mainIniPath = "/conf/activiti.properties"

func init() {
	configFile := common.ReadConfig(mainIniPath)
	if configFile == nil {
		logger.Error("read db config err")
		panic("err")
	}
	TXDB = new(sync.Map)
	mysqlConfig, err := configFile.GetSection("mysql")
	if err != nil {
		fmt.Println("get mysql conf error:", err)
		return
	}

	fillDns(mysqlConfig)
	// 启动时就打开数据库连接
	if err = initEngine(); err != nil {
		panic(err)
	}
	// 测试数据库连接是否 OK
	if err = GORM_DB.DB().Ping(); err != nil {
		panic(err)
	}
}

var (
	ConnectDBErr = errors.New("connect db error")
	UseDBErr     = errors.New("use db error")
)

func fillDns(mysqlConfig map[string]string) {
	DNS = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["username"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"])
}

func initEngine() error {
	var err error
	GORM_DB, err = gorm.Open("mysql", DNS)
	if err != nil {
		return err
	}
	GORM_DB.LogMode(true)
	return nil
}

func InitTXDB(db *gorm.DB) {
	TXDB.Store(runtime.GoroutineId(), db)
}

func ClearTXDB() {
	TXDB.Delete(runtime.GoroutineId())
}

func DB() *gorm.DB {
	db, ok := TXDB.Load(runtime.GoroutineId())
	if !ok {
		panic("TXDB not init")
	}
	return db.(*gorm.DB)
}
