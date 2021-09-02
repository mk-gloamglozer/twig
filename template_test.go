package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildFromTemplate_WritesToBuffer(t *testing.T) {
	helloFun := func() string {
		return "hello"
	}

	testTemplate := Template{
		Tmpl: "this is a template {{ hello }}",
		Funcs: map[string]interface{}{
			"hello": helloFun,
		},
	}

	final, err := testTemplate.Execute()
	assert.NoError(t, err)
	assert.Equal(t, final, "this is a template hello")
}
