package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

const testFile = "./afile"
const baseDir = "/base"

type testFileReader struct{}

func (*testFileReader) readFile(path string) (*File, error) {
	if path == filepath.Join(baseDir, testFile) {
		return &File{Name: "hello", Data: []byte("goodbye")}, nil
	} else {
		return nil, fmt.Errorf("No file with path %s was found", path)
	}
}

type testAppCtx struct{}

func (*testAppCtx) BaseFilePath() string { return "/base" }

func (*testAppCtx) Template() []byte { return []byte("random string") }

func TestFuncs(t *testing.T) {
	//TODO write tests for failure cases
	tests := []struct {
		tpl, expect string
		vars        interface{}
	}{{
		tpl:    `{{ toYaml . }}`,
		expect: `foo: bar`,
		vars:   map[string]interface{}{"foo": "bar"},
	}, {
		tpl:    `{{ toToml . }}`,
		expect: "foo = \"bar\"\n",
		vars:   map[string]interface{}{"foo": "bar"},
	}, {
		tpl:    `{{ toJson . }}`,
		expect: `{"foo":"bar"}`,
		vars:   map[string]interface{}{"foo": "bar"},
	}, {
		tpl:    `{{ fromYaml . }}`,
		expect: "map[hello:world]",
		vars:   `hello: world`,
	}, {
		tpl:    `{{ fromYamlArray . }}`,
		expect: "[one 2 map[name:helm]]",
		vars:   "- one\n- 2\n- name: helm\n",
	}, {
		tpl:    `{{ fromYamlArray . }}`,
		expect: "[one 2 map[name:helm]]",
		vars:   `["one", 2, { "name": "helm" }]`,
	}, {
		// Regression for https://github.com/helm/helm/issues/2271
		tpl:    `{{ toToml . }}`,
		expect: "[mast]\n  sail = \"white\"\n",
		vars:   map[string]map[string]string{"mast": {"sail": "white"}},
	}, {
		tpl:    `{{ fromYaml . }}`,
		expect: "map[Error:yaml: unmarshal errors:\n  line 1: cannot unmarshal !!seq into map[string]interface {}]",
		vars:   "- one\n- two\n",
	}, {
		tpl:    `{{ fromJson .}}`,
		expect: `map[hello:world]`,
		vars:   `{"hello":"world"}`,
	}, {
		tpl:    `{{ fromJson . }}`,
		expect: `map[Error:json: cannot unmarshal array into Go value of type map[string]interface {}]`,
		vars:   `["one", "two"]`,
	}, {
		tpl:    `{{ fromJsonArray . }}`,
		expect: `[one 2 map[name:helm]]`,
		vars:   `["one", 2, { "name": "helm" }]`,
	}, {
		tpl:    `{{ fromJsonArray . }}`,
		expect: `[json: cannot unmarshal object into Go value of type []interface {}]`,
		vars:   `{"hello": "world"}`,
	}, {
		tpl:    `{{ fromYaml . }}`,
		expect: "map[Error:yaml: unmarshal errors:\n  line 1: cannot unmarshal !!seq into map[string]interface {}]",
		vars:   `["one", "two"]`,
	}, {
		tpl:    `{{ fromYamlArray . }}`,
		expect: "[yaml: unmarshal errors:\n  line 1: cannot unmarshal !!map into []interface {}]",
		vars:   `hello: world`,
	}, {
		tpl:    "{{ fromFile . }}",
		expect: "goodbye",
		vars:   testFile,
	}}

	for _, tt := range tests {
		var b strings.Builder
		err := template.Must(template.New("test").Funcs(funcMap(&testFileReader{}, &testAppCtx{})).Parse(tt.tpl)).Execute(&b, tt.vars)
		assert.NoError(t, err)
		assert.Equal(t, tt.expect, b.String(), tt.tpl)
	}
}

func TestWhenFileIsMissing_ThenFromFileFailsReturnsError(t *testing.T) {
	tpl := "{{ fromFile . }}"
	vars := "./nonExistantFile"

	var b strings.Builder
	err := template.Must(
		template.New("test").
			Funcs(funcMap(&testFileReader{}, &testAppCtx{})).
			Parse(tpl)).Execute(&b, vars)
	if assert.Error(t, err) {
		if ferr, ok := err.(template.ExecError); ok {
			path := filepath.Join(baseDir, vars)
			assert.Equal(t, fmt.Errorf("No file with path %s was found", path), errors.Unwrap(ferr.Unwrap()))
		}
	}
}
