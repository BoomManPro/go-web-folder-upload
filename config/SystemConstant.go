package config

const ApiList = "/api/list/"

const Preview = "/preview/"

var PreviewPath = Preview[0 : len(Preview)-1]

var StorePath = "F:/nginx/"

var StorePathPrefix = ""

func InitStorePath(path string) {
	StorePath = path
	StorePathPrefix = StorePath[0 : len(StorePath)-1]
}
