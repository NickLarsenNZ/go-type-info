package object_info_test

import "testing"
import (
	oi "object-info"
	"reflect"
)

type customType interface {}

type S struct {
	F1 string
	F2 uint64
	f3 bool
	f4 []string
	f5 customType
}

func TestFields(t *testing.T) {
	result := oi.StructFields(&S{})

	expected := &oi.Structure{
		Fields: []oi.NameType{
			{
				Name: "F1",
				Type: "string",
			},
			{
				Name: "F2",
				Type: "uint64",
			},
			{
				Name: "f3",
				Type: "bool",
			},
			{
				Name: "f4",
				Type: "[]string",
			},
			{
				Name: "f5",
				Type: "object_info_test.customType",
			},
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %s, got: %s", expected, result)
	}
}
