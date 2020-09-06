package mysql

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB  //创建数据库 初始化

func init(){
	db,_=sql.Open("mysql","root:123456@tcp(127.0.0.1:3339)/fileserver?charset=utf8")
	db.SetMaxOpenConns(100)
	err:=db.Ping()
	if err!= nil{
		log.Fatal("failed to connect to mysql",err)
		os.Exit(1)
	}
}
//返回数据库连接
func DBConn() *sql.DB{
	return db
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range scanArgs {
			if col != nil {
				record[columns[i]] = values[i]
			}
		}
		records = append(records, record)
	}
	return records
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
