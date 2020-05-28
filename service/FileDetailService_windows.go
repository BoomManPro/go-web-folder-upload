package service

import (
	"os"
	"strconv"
	"syscall"
	"web-folder-upload/models"
)

func getFileDetail(file os.FileInfo) models.FileDetailVo {

	wFileSys := file.Sys().(*syscall.Win32FileAttributeData)
	cNanSeconds := wFileSys.CreationTime.Nanoseconds() /// 返回的是纳秒
	cSec := cNanSeconds / 1e9                          ///秒

	mNanSeconds := wFileSys.LastWriteTime.Nanoseconds() /// 返回的是纳秒
	mSec := mNanSeconds / 1e9                           ///秒
	return models.FileDetailVo{
		FileName:       file.Name(),
		FileType:       convertFileType(file.IsDir()),
		Size:           strconv.FormatInt(file.Size(), 10),
		CreateTime:     unixTime2String(cSec),
		LastModifyTime: unixTime2String(mSec),
	}

}
