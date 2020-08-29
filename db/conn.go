package db

import (
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lios/go-activiti/runtime"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

//var MasterDB *xorm.Engine

var GORM_DB *gorm.DB

var TXDB *sync.Map

var dns string
var ROOT string

const mainIniPath = "/conf/activiti.properties"

func init() {
	curFilename := os.Args[0]
	binaryPath, err := exec.LookPath(curFilename)
	if err != nil {
		panic(err)
	}

	binaryPath, err = filepath.Abs(binaryPath)
	if err != nil {
		panic(err)
	}

	ROOT = filepath.Dir(filepath.Dir(binaryPath))

	configPath := ROOT + mainIniPath

	if !fileExist(configPath) {
		curDir, _ := os.Getwd()
		pos := strings.LastIndex(curDir, "src")
		if pos == -1 {
			// panic("can't find " + mainIniPath)
			fmt.Println("can't find " + mainIniPath)
		} else {
			ROOT = curDir[:pos]

			configPath = ROOT + mainIniPath
		}
	}
	configFile, err := goconfig.LoadConfigFile(configPath)
	if err != nil {
		panic(err)
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
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["username"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"])
}

func initEngine() error {
	var err error
	GORM_DB, err = gorm.Open("mysql", dns)
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
		panic("db not init")
	}
	return db.(*gorm.DB)
}
func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
