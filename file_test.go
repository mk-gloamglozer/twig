package main

import (
	"io/ioutil"
	"testing"
)

func TestFileReadGivenPath(t *testing.T) {
	fp := "./test/testfile"

	file := readFile(fp)

	if file.Name != fp {
		t.Errorf("testfile name should have been %s but was %s", fp, file.Name)
	}

	testbytes := testRead(fp)

	for i := range file.Data {
		if file.Data[i] != testbytes[i] {
			t.Errorf("test file was not read correctly. \n\tExpected: %s \n\t     Got: %s", string(file.Data), string(testbytes))
			break
		}
	}
}

func testRead(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
