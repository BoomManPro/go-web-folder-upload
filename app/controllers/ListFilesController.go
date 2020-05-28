package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"web-folder-upload/config"
	"web-folder-upload/models"
	"web-folder-upload/service"
	"web-folder-upload/utils"
)

//上传文件handler
func ListFiles(w http.ResponseWriter, r *http.Request) {
	//获取请求地址 path

	uri, e := url.QueryUnescape(r.RequestURI)
	if e != nil {
		fmt.Print(e)
		return
	}
	path := utils.Substr(uri, len(config.ApiList)-1, len(uri))
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	vo := models.ResultVo{
		Code:     20000,
		ShowMsg:  "SUCCESS",
		ErrorMsg: "",
		Data:     service.GetPreview(path),
	}
	result, _ := json.Marshal(vo)
	io.WriteString(w, string(result))

}
