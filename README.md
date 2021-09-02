# helm-value-compiler
Compiles files containing go template instructions into helm value.yaml files

## Context

This project was created to allow repos making use of helm value files to include go style templating, 
and for this template to be compiled into a values.yaml file. 

The functions availible include all basic go template functions as well as all [sprig](https://github.com/Masterminds/sprig) 
template functions. 

## Usage

``` 
$ twig [-out <output-file>] <input-file>
```

output file will default to values.yaml in cwd. 

## Functions

all but one functions are taken from [sprig](https://github.com/Masterminds/sprig)

### fromFile

Use to load the contents of the a file into the template.

```
{{ fromFile "./path/to/file" }}
```

path is relative to the template file.
