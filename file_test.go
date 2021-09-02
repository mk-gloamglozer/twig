package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileReadGivenPath(t *testing.T) {
	fp := "./test/testfile"

	file, err := new(IFileReader).readFile(fp)

	assert.NoError(t, err, "An error occured when reading test file")

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
