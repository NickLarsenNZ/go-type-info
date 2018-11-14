package type_info

import (
	"reflect"
)

type Structure struct {
	Fields []NameType
}

type NameType struct {
	Name string
	Type string
}

func StructFields(Struct interface{}) *Structure {
	val := reflect.Indirect(reflect.ValueOf(Struct))

	var fields []NameType
	for i := 0; i < val.Type().NumField(); i++ {
		fields = append(fields, NameType{
			Name: val.Type().Field(i).Name,
			Type: val.Type().Field(i).Type.String(),
		})
	}

	return &Structure{
		Fields: fields,
	}
}
