package pkg

type ITemplate interface {
	GetTemplate() string
	GetPath() string
	GetFileName() string
	ParseData(file string)
	GetData() interface{}
	SetBaseDir(dir string)
	GetBaseDir() string
}
