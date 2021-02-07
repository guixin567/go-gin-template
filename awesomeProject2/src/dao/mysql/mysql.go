package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Xdb *sqlx.DB

func Init() (err error) {
	//"用户名:密码@tcp(地址:端口)/数据库名"
	dataSourceName := "root:12345678@tcp(127.0.0.1:3306)/zhengyinuo"
	Xdb, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return
	}
	Xdb.SetMaxOpenConns(10)
	Xdb.SetMaxIdleConns(20)
	defer Xdb.Close()
	return
}
