package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TestModel holds the schema definition for the TestModel entity.
type TestModel struct {
	ent.Schema
}

// Annotations .
func (TestModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "test_model"},
	}
}

// Fields of the TestModel.
func (TestModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("test_string"),
		field.Bool("test_bool"),
		field.Float64("test_double"),
		field.Float32("test_float"),
		field.Int32("test_int32"),
		field.Uint32("test_uint32"),
		field.Int64("test_int64"),
		field.Uint64("test_uint64"),
		field.Int32("test_sint32"),
		field.Int64("test_sint64"),
		field.Uint32("test_sfixd32"),
		field.Uint64("test_sfix64"),
		field.String("test_bytes"),
	}
}

// Edges of the TestModel.
func (TestModel) Edges() []ent.Edge {
	return []ent.Edge{}
}
