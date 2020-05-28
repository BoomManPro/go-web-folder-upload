package service

import (
	"os"
	"strconv"
	"syscall"
	"web-folder-upload/models"
)

func getFileDetail(file os.FileInfo) models.FileDetailVo {
	stat_t := file.Sys().(*syscall.Stat_t)
	cSec := int64(stat_t.Ctim.Sec)
	mSec := int64(stat_t.Mtim.Sec) / 1e9 ///秒
	return models.FileDetailVo{
		FileName:       file.Name(),
		FileType:       convertFileType(file.IsDir()),
		Size:           strconv.FormatInt(file.Size(), 10),
		CreateTime:     unixTime2String(cSec),
		LastModifyTime: unixTime2String(mSec),
	}
}
