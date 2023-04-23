package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	fmt.Println("sql init")
	dsn := "root:fate_deV2999@tcp(172.16.83.22:3306)/dataBase_zhulin?charset=utf8"
	SqlDB, err = sql.Open("mysql", dsn) //连接数据库，open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	// fmt.Println(SqlDB)
	err = SqlDB.Ping() //连接数据库，校验用户名和密码
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功~")

	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

}

func NewDb() *sql.DB {
	if SqlDB != nil {
		fmt.Println("db nil")
		return nil
	}
	return SqlDB
}

func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		fmt.Println("close stmt connection")
		stmt.Close()
	}
}
