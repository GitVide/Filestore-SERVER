package meta

import (
	mydb "filestore-server/db"
)
// FileMeta 文件元信息结构
type FileMeta struct {
	//文件的唯一标志
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas=make(map[string]FileMeta)
}

// UpdateFileMeta 新增/更新文件元信息
func UpdateFileMeta(fMeta FileMeta) {
	fileMetas[fMeta.FileSha1] = fMeta
}

// GetFileMeta 通过fileSha1获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta{
	return fileMetas[fileSha1]
}

// GetFileMetaDB 从mysql获取文件元信息
func GetFileMetaDB(filesha1 string) (FileMeta, error) {
	tfile,err:=mydb.GetFileMeta(filesha1)
	if err!=nil {
		return FileMeta{}, err
	}
	fmeta:=FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, err
}

// UpdateFileMetaDB : 新增/更新文件元信息到mysql中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(
		fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// RemoveFileMeta 删除元信息
func RemoveFileMeta(fileSha1 string)  {
	delete(fileMetas,fileSha1)
}
