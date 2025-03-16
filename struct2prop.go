package struct2prop

import (
	"errors"
	"reflect"
	"unicode"
)

type Prop struct {
	Type        PropType        `json:"type"`
	Description string          `json:"description,omitempty"`
	Items       *Prop           `json:"items,omitempty"`
	Enum        []any           `json:"enum,omitempty"`
	Properties  map[string]Prop `json:"properties,omitempty"`
	Required    []string        `json:"required,omitempty"`
}

func GetProperties(s any) (*Prop, error) {
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Struct {
		return nil, errors.New("expected a struct")
	}

	prop, err := getPropsFromType(t)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func getPropsFromType(t reflect.Type) (*Prop, error) {
	schemaType, err := getPropTypeFromKind(t.Kind())
	if err != nil {
		return nil, err
	}

	var prop Prop
	prop.Type = schemaType
	// handle object
	if schemaType == ObjectPropType {
		prop.Properties = make(map[string]Prop)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldProp, err := getPropsFromType(field.Type)
			if err != nil {
				return nil, err
			}
			description := field.Tag.Get("description")
			if description != "" {
				fieldProp.Description = description
			}
			r := []rune(field.Name)
			r[0] = unicode.ToLower(r[0])
			prop.Properties[string(r)] = *fieldProp
		}
		return &prop, nil
	}
	// handle array
	if schemaType == ArrayPropType {
		fieldProp, err := getPropsFromType(t.Elem())
		if err != nil {
			return nil, err
		}
		prop.Items = fieldProp
		return &prop, nil
	}

	return &prop, nil
}

type PropType string

const (
	StringPropType PropType = "string"
	IntPropType    PropType = "integer"
	NumberPropType PropType = "number"
	BoolPropType   PropType = "boolean"
	ObjectPropType PropType = "object"
	ArrayPropType  PropType = "array"
)

func getPropTypeFromKind(k reflect.Kind) (PropType, error) {
	// check if string
	if k == reflect.String {
		return StringPropType, nil
	}
	// check if an Integer
	if k >= reflect.Int && k <= reflect.Uint64 {
		return IntPropType, nil
	}
	// check if float
	if k == reflect.Float32 || k == reflect.Float64 {
		return NumberPropType, nil
	}
	// check if bool
	if k == reflect.Bool {
		return BoolPropType, nil
	}
	// check if struct
	if k == reflect.Struct {
		return ObjectPropType, nil
	}
	if k == reflect.Slice {
		return ArrayPropType, nil
	}

	return "", nil
}
