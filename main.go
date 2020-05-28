package main

import (
	"fmt"
	"net/http"
	"web-folder-upload/app/controllers"
	"web-folder-upload/config"
)

func main() {
	// preview
	http.Handle(config.Preview, http.StripPrefix(config.Preview, http.FileServer(http.Dir(config.StorePath))))
	//前端静态界面接口
	http.Handle("/", http.FileServer(http.Dir("./html")))
	//文件列表接口
	http.HandleFunc(config.ApiList, controllers.ListFiles)
	//上传文件接口
	http.HandleFunc("/api/uploadDirectory", controllers.UploadFolderHandler)
	//上传文件夹接口
	e := http.ListenAndServe(":8080", nil)
	fmt.Println(e)
}
