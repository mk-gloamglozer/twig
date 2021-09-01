package main

type files map[string][]byte

type Context struct {
	Files files
}

func newContext(files []*File) *Context {
	return &Context{
		Files: createFiles(files),
	}
}

func createFiles(from []*File) files {
	files := make(map[string][]byte)
	for _, f := range from {
		files[f.Name] = f.Data
	}
	return files
}
