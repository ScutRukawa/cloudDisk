package meta

import (
	"cloudDisk/db"
	"log"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
	log.Println("update a file mata data")
}

func UpdateFileMetaDb(fmeta FileMeta) bool {
	ret := db.OnFileUploadFinish(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
	log.Println("OnFileUploadFinish success")
	if ret != true {
		log.Fatal("update file to DB failed")
		return false
	}
	log.Println("OnFileUploadFinish success")
	return true
}
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//从db 获取文件元信息
func GetFileMetaDB(fileSha1 string) (*FileMeta, error) {
	tfile, err := db.GetFileMeta(fileSha1)
	if err != nil {
		return &FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return &fmeta, nil
}
