package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"web-folder-upload/config"
	"web-folder-upload/models"
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

		if files != nil {
			err = uploadFolder(path, files)
		}
		file := m.File["file"]
		if file != nil {
			err = uploadFile(path, file[0])
		}

		if err != nil {
			result, _ := json.Marshal(models.ResultVo{
				Code:     50000,
				ShowMsg:  "Error",
				ErrorMsg: err.Error(),
				Data:     nil,
			})
			fmt.Fprint(w, string(result))

			return
		}

		result, _ := json.Marshal(models.ResultVo{
			Code:     20000,
			ShowMsg:  "SUCCESS",
			ErrorMsg: "",
			Data:     nil,
		})

		fmt.Fprint(w, string(result))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func uploadFile(path string, file0 *multipart.FileHeader) error {
	//StorePath 去除末尾 /
	filePath := config.StorePath + path + file0.Filename
	//for each fileheader, get a handle to the actual file
	file, err := file0.Open()
	defer file.Close()
	if err != nil {
		return err
	}
	//create destination file making sure the path is writeable.
	createFile(filePath)
	dst, err := os.Create(filePath)
	defer dst.Close()
	if err != nil {
		return err
	}
	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	return nil
}

func uploadFolder(path string, files []*multipart.FileHeader) error {
	for i, _ := range files {
		err := uploadFile(path, files[i])
		if err != nil {
			return err
		}

	}
	return nil
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	dir := filePath[0:strings.LastIndex(filePath, "/")]
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
