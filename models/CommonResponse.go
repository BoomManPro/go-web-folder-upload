package models

type ResultVo struct {

	//code
	Code int `json:"code"`

	//show ErrorMsg
	ShowMsg string `json:"showMsg"`

	//错误信息详情
	ErrorMsg string `json:"errorMsg"`

	//返回数据
	Data interface{} `json:"data"`
}
