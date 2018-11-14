package type_info_test

import "testing"
import (
	"reflect"
	"go-type-info" // Todo: Not sure how to make this not bind to my local directory name. If I add the full github URL, then my coverage shows as 0% with tests passing anyway.
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
	result := type_info.StructFields(&S{})

	expected := &type_info.Structure{
		Fields: []type_info.NameType{
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
