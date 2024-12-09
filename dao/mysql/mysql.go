package mysql

import (
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init() (err error) {
	// "user:password@tcp(host:port)/dbname"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
	//	viper.GetString("mysql.user"),
	//	viper.GetString("mysql.password"),
	//	viper.GetInt("mysql.host"),
	//	viper.GetInt("mysql.port"),
	//	viper.GetString("mysql.dbname"))

	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), // 使用 GetString 而不是 GetInt
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"))
	db, err = sqlx.Connect("mysql", dsn2)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("max_idle_conns"))
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
