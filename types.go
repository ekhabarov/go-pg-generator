package main

import (
	"fmt"
	"log"
)

const (
	tpl = `type %s struct {
%s}
`
)

var (
	// key - pg type, value - go type
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

func getStruct(tab string, cols []*column) string {
	var body string
	for _, c := range cols {
		body += fmt.Sprintf("\t%s %s `db:%q json:%[4]q` // sqltype: %s\n",
			snake2Camel(c.Name), convertType(c), c.Name, snake2CamelLower(c.Name), c.Type)
	}

	return fmt.Sprintf(tpl, snake2Camel(tab), body)
}

func getMap(n bool) map[string]string {
	if n {
		return nullableTypes
	}
	return nonNullableTypes
}

func convertType(c *column) string {
	t, ok := getMap(c.IsNull)[c.Type]
	if !ok {
		log.Fatalln("unknown type: ", t)
	}

	if c.IsArray {
		return "[]" + t
	}
	return t
}
