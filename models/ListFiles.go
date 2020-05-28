package models

//文件类型
type FileType string

const (
	//文件夹
	DIRECTORY FileType = "DIRECTORY"
	//文件
	FILE FileType = "FILE"
)

type FileDetailVo struct {
	//文件名
	FileName string `json:"fileName"`

	//文件类型
	FileType FileType `json:"fileType"`

	//文件大小
	Size string `json:"size"`

	//创建时间
	CreateTime string `json:"createTime"`

	//最后修改时间
	LastModifyTime string `json:"lastModifyTime"`
}

type PreviewVo struct {

	//预览目录
	PreviewPath string `json:"previewPath"`

	//文件信息
	Files []FileDetailVo `json:"files"`
}
