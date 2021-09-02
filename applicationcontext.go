package main

type ApplicationContext interface {
	BaseFilePath() string
	Template() []byte
}

type IApplicationContext struct {
	BaseFile *File
}

func (ctx *IApplicationContext) BaseFilePath() string {
	return ctx.BaseFile.Name
}

func (ctx *IApplicationContext) Template() []byte {
	return ctx.BaseFile.Data
}

func newApplicationContext(file *File) ApplicationContext {
	return &IApplicationContext{BaseFile: file}
}
