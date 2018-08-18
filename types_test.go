package main

import "testing"

type TestColumn struct {
	col    *column
	result string
}

func NewTC(t string, in, ia bool, r string) *TestColumn {
	return &TestColumn{
		&column{
			Type:    t,
			IsNull:  in,
			IsArray: ia},
		r,
	}
}
func TestCovertTypes(t *testing.T) {
	columns := [...]*TestColumn{
		//Non-null types
		NewTC("int2", false, false, "int16"),
		NewTC("int4", false, false, "int32"),
		NewTC("int8", false, false, "int64"),
		NewTC("float4", false, false, "float32"),
		NewTC("float8", false, false, "float64"),
		NewTC("numeric", false, false, "float64"),
		NewTC("money", false, false, "float64"),
		NewTC("bpchar", false, false, "string"),
		NewTC("varchar", false, false, "string"),
		NewTC("text", false, false, "string"),
		NewTC("bytea", false, false, "[]byte"),
		NewTC("uuid", false, false, "uuid.UUID"),
		NewTC("timestamp", false, false, "time.Time"),
		NewTC("timestamptz", false, false, "time.Time"),
		NewTC("time", false, false, "time.Time"),
		NewTC("timetz", false, false, "time.Time"),
		NewTC("date", false, false, "time.Time"),
		NewTC("interval", false, false, "time.Time"),
		NewTC("bool", false, false, "bool"),
		NewTC("bit", false, false, "uint32"),
		NewTC("varbit", false, false, "uint32"),
		NewTC("json", false, false, "struct{...}"),
		NewTC("xml", false, false, "struct{...}"),

		//Nullable types
		NewTC("int2", true, false, "null.Int"),
		NewTC("int4", true, false, "null.Int"),
		NewTC("int8", true, false, "null.Int"),
		NewTC("float4", true, false, "null.Float"),
		NewTC("float8", true, false, "null.Float"),
		NewTC("numeric", true, false, "null.Float"),
		NewTC("money", true, false, "null.Float"),
		NewTC("bpchar", true, false, "null.String"),
		NewTC("varchar", true, false, "null.String"),
		NewTC("text", true, false, "null.String"),
		NewTC("bytea", true, false, "*[]byte"),
		NewTC("uuid", true, false, "*uuid.UUID"),
		NewTC("timestamp", true, false, "null.Time"),
		NewTC("timestamptz", true, false, "null.Time"),
		NewTC("time", true, false, "null.Time"),
		NewTC("timetz", true, false, "null.Time"),
		NewTC("date", true, false, "null.Time"),
		NewTC("interval", true, false, "null.Time"),
		NewTC("bool", true, false, "null.Bool"),
		NewTC("bit", true, false, "*uint32"),
		NewTC("varbit", true, false, "*uint32"),
		NewTC("json", true, false, "*struct{...}"),
		NewTC("xml", true, false, "*struct{...}"),

		//Non-null array types
		NewTC("int2", false, true, "[]int16"),
		NewTC("int4", false, true, "[]int32"),
		NewTC("int8", false, true, "[]int64"),
		NewTC("float4", false, true, "[]float32"),
		NewTC("float8", false, true, "[]float64"),
		NewTC("numeric", false, true, "[]float64"),
		NewTC("money", false, true, "[]float64"),
		NewTC("bpchar", false, true, "[]string"),
		NewTC("varchar", false, true, "[]string"),
		NewTC("text", false, true, "[]string"),
		NewTC("bytea", false, true, "[][]byte"),
		NewTC("uuid", false, true, "[]uuid.UUID"),
		NewTC("timestamp", false, true, "[]time.Time"),
		NewTC("timestamptz", false, true, "[]time.Time"),
		NewTC("time", false, true, "[]time.Time"),
		NewTC("timetz", false, true, "[]time.Time"),
		NewTC("date", false, true, "[]time.Time"),
		NewTC("interval", false, true, "[]time.Time"),
		NewTC("bool", false, true, "[]bool"),
		NewTC("bit", false, true, "[]uint32"),
		NewTC("varbit", false, true, "[]uint32"),
		NewTC("json", false, true, "[]struct{...}"),
		NewTC("xml", false, true, "[]struct{...}"),

		//Nullable array types
		NewTC("int2", true, true, "[]null.Int"),
		NewTC("int4", true, true, "[]null.Int"),
		NewTC("int8", true, true, "[]null.Int"),
		NewTC("float4", true, true, "[]null.Float"),
		NewTC("float8", true, true, "[]null.Float"),
		NewTC("numeric", true, true, "[]null.Float"),
		NewTC("money", true, true, "[]null.Float"),
		NewTC("bpchar", true, true, "[]null.String"),
		NewTC("varchar", true, true, "[]null.String"),
		NewTC("text", true, true, "[]null.String"),
		NewTC("bytea", true, true, "[]*[]byte"),
		NewTC("uuid", true, true, "[]*uuid.UUID"),
		NewTC("timestamp", true, true, "[]null.Time"),
		NewTC("timestamptz", true, true, "[]null.Time"),
		NewTC("time", true, true, "[]null.Time"),
		NewTC("timetz", true, true, "[]null.Time"),
		NewTC("date", true, true, "[]null.Time"),
		NewTC("interval", true, true, "[]null.Time"),
		NewTC("bool", true, true, "[]null.Bool"),
		NewTC("bit", true, true, "[]*uint32"),
		NewTC("varbit", true, true, "[]*uint32"),
		NewTC("json", true, true, "[]*struct{...}"),
		NewTC("xml", true, true, "[]*struct{...}"),
	}

	for _, c := range columns {
		if convertType(c.col) != c.result {
			t.Errorf("invalid typecast: pg type: name: %s, isnull: %t, isarray: %t  => go type: %s", c.col.Type, c.col.IsNull, c.col.IsArray, c.result)
		}
	}

}
