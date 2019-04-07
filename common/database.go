package common

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Database struct {
	*gorm.DB
}

var (
	MySQL *gorm.DB
)

func InitMySQL() *gorm.DB {

	var (
		db *gorm.DB
		e  error
	)

	mysqlConnectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True&loc=%s", viper.GetString("storage.mysql.user"), viper.GetString("storage.mysql.password"), viper.GetString("storage.mysql.host"), viper.GetString("storage.mysql.port"), viper.GetString("storage.mysql.database"), viper.GetString("storage.mysql.timezone"))
	// Reconnect
	for db, e = gorm.Open("mysql", mysqlConnectString); e != nil; {
		fmt.Println("Gorm Open DB Err: ", e)
		log.Println(fmt.Sprintf("GORM cannot connect to database, retry in %d seconds...", viper.GetInt("storage.mysql.retry_interval")))
		time.Sleep(time.Duration(viper.GetInt("storage.mysql.retry_interval")) * time.Second)
	}

	log.Println("Connected to database ", viper.GetString("storage.mysql.user"), viper.GetString("storage.mysql.password"), viper.GetString("storage.mysql.host")+":"+viper.GetString("storage.mysql.port"), viper.GetString("storage.mysql.database"))
	db.DB().SetMaxIdleConns(viper.GetInt("storage.mysql.max_idle_conns"))
	MySQL = db
	return MySQL
}

func GetMySQL() *gorm.DB {
	return MySQL
}
