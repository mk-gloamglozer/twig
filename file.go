package main

import (
	"io/ioutil"
	"log"
)

type FileReader interface {
	readFile(string) (*File, error)
}

type IFileReader struct {
}

type File struct {
	Name string
	Data []byte
}

func defaultReader() FileReader {
	return &IFileReader{}
}

func (*IFileReader) readFile(path string) (*File, error) {

	var data, err = ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &File{
		Name: path,
		Data: data,
	}, nil
}
