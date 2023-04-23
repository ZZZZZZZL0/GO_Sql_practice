package service

import (
	"fmt"

	"github.com/duyanghao/gin-apiserver/pkg/db"
	"github.com/duyanghao/gin-apiserver/pkg/models"

	_ "github.com/go-sql-driver/mysql" // _ 代表导入包但不使用，init()
)

type user_student struct {
	student_id   int
	student_name string
	student_age  int
}

func init() {
	//to do client 赋值
}

type ToDoService interface {
	Get(string) int
	// Get(studentName) return studentID
}

type toDoService struct {
}

func NewToDoService() ToDoService {
	return &toDoService{}
}

// 按照name查询，返回num
// toDoService结构体实现了ToDoService接口，返回值是string
func (c *toDoService) Get(testInput string) int {
	// 单行查询，按name
	/*
		dsn := "root:fate_deV2999@tcp(172.16.83.22:3306)/dataBase_zhulin"
		db, err := sql.Open("mysql", dsn) //连接数据库，open不会检验用户名和密码
		if err != nil {
			fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
			return -1
		}
		err = db.Ping() //连接数据库，校验用户名和密码
		if err != nil {
			fmt.Printf("open %s faild,err:%v\n", dsn, err)
			return -1
		}
		fmt.Println("连接数据库成功~")
	*/
	sqlStr := "select Num,Name from TestTable where Name=?;"
	//2.执行
	name_query := testInput
	rowObj := db.SqlDB.QueryRow(sqlStr, name_query) //从连接池里取一个连接出来去数据库查询单挑记录
	//3.拿到结果
	var u1 models.StudentClass1
	rowObj.Scan(&u1.StudentID, &u1.StudentName, &u1.StudentAge)
	//打印结果
	fmt.Printf("u1:%#v\n", u1)
	fmt.Println("----------------------above: single query")
	return u1.StudentID
}
