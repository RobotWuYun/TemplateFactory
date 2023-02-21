package constants

var PbField2EntMap = map[string]string{
	"string":   "String",
	"bool":     "Bool",
	"double":   "Float64",
	"float":    "Float32",
	"int32":    "Int32",
	"uint32":   "Uint32",
	"int64":    "Int64",
	"uint64":   "Uint64",
	"sint32":   "Int32",
	"sint64":   "Int64",
	"sfixed32": "Uint32",
	"sfixed64": "Uint64",
	"bytes":    "String",
}

var PbField2StructMap = map[string]string{
	"string":   "string",
	"bool":     "bool",
	"double":   "float64",
	"float":    "float32",
	"int32":    "int32",
	"uint32":   "uint32",
	"int64":    "int64",
	"uint64":   "uint64",
	"sint32":   "int32",
	"sint64":   "int64",
	"sfixed32": "uint32",
	"sfixed64": "uint64",
	"bytes":    "string",
}

var PbField2SqlMap = map[string]string{
	"string":   "VARCHAR()",
	"bool":     "TINYINT(1)",
	"double":   "DOUBLE",
	"float":    "FLOAT",
	"int32":    "INTEGER",
	"uint32":   "uint32",
	"int64":    "BIGINT",
	"uint64":   "uint64",
	"sint32":   "INTEGER",
	"sint64":   "BIGINT",
	"sfixed32": "uint32",
	"sfixed64": "uint64",
	"bytes":    "VARCHAR()",
}
