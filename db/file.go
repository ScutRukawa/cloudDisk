package db

import (
	mydb "cloudDisk/db/mysql"
	"database/sql"
	"fmt"
	"log"
)

//文件上传完成，保存meta
func OnFileUploadFinish(filehash string, filename string, filesize int64, fileaddr string) bool {
	log.Printf("OnFileUploadFinish success %s %s %d %s", filehash, filename, filesize, fileaddr)
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`," +
		"`file_addr`,`status`) values(?,?,?,?,1)") //防止注入攻击，
	if err != nil {
		fmt.Println("Prepare failed", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println("Exec failed", err)
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("file with hash: %s has been upload before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString //??
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//从db 获取文件元信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file " +
			"where file_sha1=? and status=1 limit 1")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	title := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&title.FileHash, &title.FileAddr, &title.FileName, &title.FileSize)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &title, err
}
