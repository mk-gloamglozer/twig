package main

import (
	"io/ioutil"
	"log"
)

var _readFile = ioutil.ReadFile

type File struct {
	Name string
	Data []byte
}

func readFile(path string) *File {

	var data, err = _readFile(path)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return &File{
		Name: path,
		Data: data,
	}
}
