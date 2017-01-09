package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/inflection"
)

const (
	tpl = `package mrog

type %s struct {
%s}`
)

var (
	//key - pg type, value - go type
	nonNullableTypes map[string]string = map[string]string{
		"int2":        "int16",
		"int4":        "int32",
		"int8":        "int64",
		"float4":      "float32",
		"float8":      "float64",
		"numeric":     "float64",
		"money":       "float64",
		"bpchar":      "string",
		"varchar":     "string",
		"text":        "string",
		"bytea":       "[]byte",
		"uuid":        "uuid.UUID",
		"timestamp":   "time.Time",
		"timestamptz": "time.Time",
		"time":        "time.Time",
		"timetz":      "time.Time",
		"date":        "time.Time",
		"interval":    "time.Time",
		"bool":        "bool",
		"bit":         "uint32",
		"varbit":      "uint32",
		"json":        "struct{...}",
		"xml":         "struct{...}",
	}

	nullableTypes map[string]string = map[string]string{
		"int2":        "null.Int",
		"int4":        "null.Int",
		"int8":        "null.Int",
		"float4":      "null.Float",
		"float8":      "null.Float",
		"numeric":     "null.Float",
		"money":       "null.Float",
		"bpchar":      "null.String",
		"varchar":     "null.String",
		"text":        "null.String",
		"bytea":       "*[]byte",
		"uuid":        "*uuid.UUID",
		"timestamp":   "null.Time",
		"timestamptz": "null.Time",
		"time":        "null.Time",
		"timetz":      "null.Time",
		"date":        "null.Time",
		"interval":    "null.Time",
		"bool":        "null.Bool",
		"bit":         "*uint32",
		"varbit":      "*uint32",
		"json":        "*struct{...}",
		"xml":         "*struct{...}",
	}
)

func GetStruct(tab string, cols []*Column) (title string, s string) {
	var body string
	title = inflection.Singular(snake2Camel(tab))
	for _, c := range cols {
		body += fmt.Sprintf("\t%s %s\n", snake2Camel(c.Name), pg2GoType(c))
	}

	return title, fmt.Sprintf(tpl, title, body)
}

func getMap(n bool) map[string]string {
	if n {
		return nullableTypes
	} else {
		return nonNullableTypes
	}
}

func pg2GoType(c *Column) (t string) {
	if f, ok := getMap(c.IsNull)[c.Type]; ok {
		t = f
	} else {
		log.Fatalf("unknown type: ", f)
	}

	if c.IsArray {
		return "[]" + t
	} else {
		return t
	}
}
