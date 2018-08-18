package main

import "testing"

type TestColumn struct {
	col    *column
	result string
}

func newTC(t string, in, ia bool, r string) *TestColumn {
	return &TestColumn{
		&column{
			Type:    t,
			IsNull:  in,
			IsArray: ia},
		r,
	}
}
func TestConvertTypes(t *testing.T) {
	columns := [...]*TestColumn{
		//Non-null types
		newTC("int2", false, false, "int16"),
		newTC("int4", false, false, "int32"),
		newTC("int8", false, false, "int64"),
		newTC("float4", false, false, "float32"),
		newTC("float8", false, false, "float64"),
		newTC("numeric", false, false, "float64"),
		newTC("money", false, false, "float64"),
		newTC("bpchar", false, false, "string"),
		newTC("varchar", false, false, "string"),
		newTC("text", false, false, "string"),
		newTC("bytea", false, false, "[]byte"),
		newTC("uuid", false, false, "uuid.UUID"),
		newTC("timestamp", false, false, "time.Time"),
		newTC("timestamptz", false, false, "time.Time"),
		newTC("time", false, false, "time.Time"),
		newTC("timetz", false, false, "time.Time"),
		newTC("date", false, false, "time.Time"),
		newTC("interval", false, false, "time.Time"),
		newTC("bool", false, false, "bool"),
		newTC("bit", false, false, "uint32"),
		newTC("varbit", false, false, "uint32"),
		newTC("json", false, false, "struct{...}"),
		newTC("xml", false, false, "struct{...}"),

		//Nullable types
		newTC("int2", true, false, "null.Int"),
		newTC("int4", true, false, "null.Int"),
		newTC("int8", true, false, "null.Int"),
		newTC("float4", true, false, "null.Float"),
		newTC("float8", true, false, "null.Float"),
		newTC("numeric", true, false, "null.Float"),
		newTC("money", true, false, "null.Float"),
		newTC("bpchar", true, false, "null.String"),
		newTC("varchar", true, false, "null.String"),
		newTC("text", true, false, "null.String"),
		newTC("bytea", true, false, "*[]byte"),
		newTC("uuid", true, false, "*uuid.UUID"),
		newTC("timestamp", true, false, "null.Time"),
		newTC("timestamptz", true, false, "null.Time"),
		newTC("time", true, false, "null.Time"),
		newTC("timetz", true, false, "null.Time"),
		newTC("date", true, false, "null.Time"),
		newTC("interval", true, false, "null.Time"),
		newTC("bool", true, false, "null.Bool"),
		newTC("bit", true, false, "*uint32"),
		newTC("varbit", true, false, "*uint32"),
		newTC("json", true, false, "*struct{...}"),
		newTC("xml", true, false, "*struct{...}"),

		//Non-null array types
		newTC("int2", false, true, "[]int16"),
		newTC("int4", false, true, "[]int32"),
		newTC("int8", false, true, "[]int64"),
		newTC("float4", false, true, "[]float32"),
		newTC("float8", false, true, "[]float64"),
		newTC("numeric", false, true, "[]float64"),
		newTC("money", false, true, "[]float64"),
		newTC("bpchar", false, true, "[]string"),
		newTC("varchar", false, true, "[]string"),
		newTC("text", false, true, "[]string"),
		newTC("bytea", false, true, "[][]byte"),
		newTC("uuid", false, true, "[]uuid.UUID"),
		newTC("timestamp", false, true, "[]time.Time"),
		newTC("timestamptz", false, true, "[]time.Time"),
		newTC("time", false, true, "[]time.Time"),
		newTC("timetz", false, true, "[]time.Time"),
		newTC("date", false, true, "[]time.Time"),
		newTC("interval", false, true, "[]time.Time"),
		newTC("bool", false, true, "[]bool"),
		newTC("bit", false, true, "[]uint32"),
		newTC("varbit", false, true, "[]uint32"),
		newTC("json", false, true, "[]struct{...}"),
		newTC("xml", false, true, "[]struct{...}"),

		//Nullable array types
		newTC("int2", true, true, "[]null.Int"),
		newTC("int4", true, true, "[]null.Int"),
		newTC("int8", true, true, "[]null.Int"),
		newTC("float4", true, true, "[]null.Float"),
		newTC("float8", true, true, "[]null.Float"),
		newTC("numeric", true, true, "[]null.Float"),
		newTC("money", true, true, "[]null.Float"),
		newTC("bpchar", true, true, "[]null.String"),
		newTC("varchar", true, true, "[]null.String"),
		newTC("text", true, true, "[]null.String"),
		newTC("bytea", true, true, "[]*[]byte"),
		newTC("uuid", true, true, "[]*uuid.UUID"),
		newTC("timestamp", true, true, "[]null.Time"),
		newTC("timestamptz", true, true, "[]null.Time"),
		newTC("time", true, true, "[]null.Time"),
		newTC("timetz", true, true, "[]null.Time"),
		newTC("date", true, true, "[]null.Time"),
		newTC("interval", true, true, "[]null.Time"),
		newTC("bool", true, true, "[]null.Bool"),
		newTC("bit", true, true, "[]*uint32"),
		newTC("varbit", true, true, "[]*uint32"),
		newTC("json", true, true, "[]*struct{...}"),
		newTC("xml", true, true, "[]*struct{...}"),
	}

	for _, c := range columns {
		if convertType(c.col) != c.result {
			t.Errorf("invalid typecast: pg type: name: %s, isnull: %t, isarray: %t  => go type: %s", c.col.Type, c.col.IsNull, c.col.IsArray, c.result)
		}
	}

}
