package database

import (
	"fmt"
	"log"
	"sync"

	config "github.com/WebDesign/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

//第一种写法
var DB *gorm.DB
var once sync.Once

/* 第二种写法
 var lock = &sync.Mutex{}

//包外不可以调用
//写个接口专门给客户端调用
func init() {
	var err error
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating single instance-DB now.")
			db, err = gorm.Open("mysql", getDBConfig())
			if err != nil {
				log.Fatal("Open database failed",
					zap.String("reason", err.Error()),
					zap.String("detail", "Database connection failed."))
			}

		} else {
			log.Fatal("Single instance already created.")
		}
	}

} */

func init() {
	var err error
	if DB == nil {
		once.Do(func() {
			fmt.Println("Creating database single instance now.")
			DB, err = gorm.Open("mysql", getDBConfig())
			if err != nil {
				log.Fatal("Open database failed",
					zap.String("reason", err.Error()),
					zap.String("detail", "Database connection failed."))

			}

		})

	} else {
		log.Fatal("Warning:connect to the database again",
			zap.String("reason", "Database instance already created."),
			zap.String("detail", "Connect to Database again failed."))
	}

}

func getDBConfig() string {
	username := config.GetGlobalConfig().GetString("mysql.user")
	password := config.GetGlobalConfig().GetString("mysql.password")
	addr := config.GetGlobalConfig().GetString("mysql.ip")
	name := config.GetGlobalConfig().GetString("mysql.database")
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)
	return config
}

func GetDBInstance() *gorm.DB {
	return DB
}
func DBInstanceClose() {
	DB.Close()
}