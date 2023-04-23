package service

import (
	"fmt"
	"log"

	"github.com/duyanghao/gin-apiserver/pkg/db"
	"github.com/duyanghao/gin-apiserver/pkg/models"

	_ "github.com/go-sql-driver/mysql" // _ 代表导入包但不使用，init()
)

func init() {
	//to do client 赋值
}

type SqlStudent_Service interface {
	GetID_st(int) models.StudentClass1
	Add_st(models.StudentClass1)
	Update_st(models.StudentClass1)
	Del_st(int) string
}

type sqlStudent_Service struct {
}

func Sql_new_Service() SqlStudent_Service {
	return &sqlStudent_Service{}
}

func (c *sqlStudent_Service) GetID_st(idQueryInput int) models.StudentClass1 {
	fmt.Println("--------------------service/database_td.go : GetID_st begins")

	sqlStr := "select id,name,age from student_class1 where id=?;"
	//2.调用查询函数，得到结果rowObj
	fmt.Println("before")
	rowObj := db.SqlDB.QueryRow(sqlStr, idQueryInput) //从连接池里取一个连接出来去数据库查询单挑记录
	//3.调用scan函数将结果映射到结构体 u1 中
	fmt.Println("after read")
	var u1 models.StudentClass1
	rowObj.Scan(&u1.StudentID, &u1.StudentName, &u1.StudentAge)

	//打印结果
	fmt.Printf("查询结果：u1:%#v\n", u1)
	fmt.Println("above: single query")
	return u1
}

func (c *sqlStudent_Service) Add_st(add_info models.StudentClass1) {
	// ---------------------------------------------
	// add_id := 207
	// add_name := "zoy"
	// add_age := 100
	add_id := add_info.StudentID
	add_name := add_info.StudentName
	add_age := add_info.StudentAge
	// ---------------------------------------------
	fmt.Println("--------------------service/database_td.go : Add_st begins")

	fmt.Printf("即将要添加的信息:%v\n", add_info)

	sqlStr := `insert into student_class1(id,name,age) values(?, ?, ?)` //sql语句
	ret, err := db.SqlDB.Exec(sqlStr, add_id, add_name, add_age)        //执行sql语句
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	//如果是插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	n, _ := ret.RowsAffected()
	fmt.Printf("--------------------数据新增成功,更新了%d行数据\n", n)
	fmt.Printf("id: %d", id)
}

func (c *sqlStudent_Service) Update_st(update_info models.StudentClass1) {
	// ---------------------------------------------
	// newage := 99
	// newname := "biber"
	// idx_id := 102
	newage := update_info.StudentAge
	newname := update_info.StudentName
	idx_id := update_info.StudentID
	// ---------------------------------------------
	fmt.Println("--------------------service/database_td.go : Update_st begins")

	sqlStr := `update student_class1 set name=?, age=? where id=?`
	ret, err := db.SqlDB.Exec(sqlStr, newname, newage, idx_id)
	if err != nil {
		fmt.Printf("update failed ,err:%v\n", err)
		return
	}
	n, _ := ret.RowsAffected()
	fmt.Printf("--------------------数据更新成功,更新了%d行数据\n", n)
}

func (c *sqlStudent_Service) Del_st(idDelInput int) string {
	// ---------------------------------------------
	// idDelInput := 207
	// ---------------------------------------------
	fmt.Println("--------------------service/database_td.go : Del_st begins")

	sql := "DELETE FROM student_class1 WHERE id = ?"
	res, err3 := db.SqlDB.Exec(sql, idDelInput)
	if err3 != nil {
		panic(err3.Error())
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if affectedRows == 0 {
		fmt.Printf("--------------------数据库中未查询到 ID= %d 的数据条目\n", idDelInput)
	} else {
		fmt.Printf("--------------------成功删除了ID= %d 的 %d 行数据\n", idDelInput, affectedRows)
	}
	return "successful DEL"
}
