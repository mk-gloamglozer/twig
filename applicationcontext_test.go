package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCtxFileReader struct {
	rf (func(string) (*File, error))
}

func (tfr *testCtxFileReader) readFile(path string) (*File, error) {
	return tfr.rf(path)
}

func Test_AppContext_BasePathDir_ReturnsOnlyDir(t *testing.T) {
	appctx := IApplicationContext{
		baseFile: &File{
			Name: "/path/to/file.tmpl",
			Data: []byte("some data"),
		},
		fileReader: &testCtxFileReader{
			rf: func(s string) (*File, error) { return nil, nil },
		},
	}

	assert.Equal(t, "/path/to", appctx.BaseFilePath())
}
