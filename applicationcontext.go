package main

type ApplicationContext interface {
	BaseFilePath() string
	Template() []byte
	FileReader() FileReader
}

type IApplicationContext struct {
	baseFile   *File
	fileReader FileReader
}

func (ctx *IApplicationContext) FileReader() FileReader {
	return ctx.fileReader
}

func (ctx *IApplicationContext) BaseFilePath() string {
	return ctx.baseFile.Name
}

func (ctx *IApplicationContext) Template() []byte {
	return ctx.baseFile.Data
}

func newApplicationContext(file *File) ApplicationContext {
	return &IApplicationContext{
		baseFile:   file,
		fileReader: defaultReader(),
	}
}
