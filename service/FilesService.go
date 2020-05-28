package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"web-folder-upload/config"
	"web-folder-upload/models"
)

func GetPreview(path string) (vo models.PreviewVo) {
	dir, e := ioutil.ReadDir(config.StorePath + path)
	if e != nil {
		fmt.Println("read dir error ")
		return
	}
	return models.PreviewVo{
		PreviewPath: config.PreviewPath,
		//path calc
		Files: getAllFiles(dir),
	}
}

func getAllFiles(dir []os.FileInfo) (files []models.FileDetailVo) {
	for _, fi := range dir {
		files = append(files, getFileDetail(fi))
	}
	return files
}

func unixTime2String(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func convertFileType(dir bool) models.FileType {
	if dir {
		return models.DIRECTORY
	}
	return models.FILE
}
