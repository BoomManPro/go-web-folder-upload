package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"web-folder-upload/config"
	"web-folder-upload/utils"
)

//上传文件handler
func UploadFolderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//path地址
		path := m.Value["path"][0]

		//get the *fileheaders
		files := m.File["folder"]

		for i, _ := range files {
			//StorePath 去除末尾 /
			filePath := config.StorePath + path + files[i].Filename
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			createFile(filePath)
			dst, err := os.Create(filePath)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	dir := utils.Substr(filePath, 0, strings.LastIndex(filePath, "/"))
	if !isExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
