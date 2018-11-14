package type_info_test

import "testing"
import (
	"reflect"
	info "github.com/nicklarsennz/go-type-info"
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
	result := info.StructFields(&S{})

	expected := &info.Structure{
		Fields: []info.NameType{
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
				Type: "type_info_test.customType",
			},
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %s, got: %s", expected, result)
	}
}
