package mysql

import (
	"awesomeProject/awesomeProject2/src/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Xdb *sqlx.DB

func Init() (err error) {
	//"用户名:密码@tcp(地址:端口)/数据库名"
	var mysqlConfig = config.GlobalConfig.MySql
	dataSourceName := fmt.Sprintf("s%:%s@tcp(s%:s%)/%s", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)
	Xdb, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return
	}
	Xdb.SetMaxOpenConns(10)
	Xdb.SetMaxIdleConns(20)
	return
}
