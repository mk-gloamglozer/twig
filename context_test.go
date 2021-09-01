package main

import (
	"testing"
)

func TestContextCreatedFromFile(t *testing.T) {
	byteArr := []byte("hi there buddy")
	file := &File{
		Name: "fileName",
		Data: byteArr,
	}

	context := newContext([]*File{file})
	if len(context.Files) != 1 {
		t.Errorf("%d files expect, got %d", 1, len(context.Files))
	}

	if context.Files["fileName"] == nil {
		t.Errorf("Expected to find file of name \"fileName\" but none was found")
	}

	b := context.Files["fileName"]
	for i := range b {
		if b[i] != byteArr[i] {
			t.Errorf("Expected file data \"%s\" but found \"%s\"", string(byteArr), string(b))
		}
	}

}
