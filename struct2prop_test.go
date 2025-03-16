package struct2prop

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPropTypeFromKind(t *testing.T) {
	var (
		String  string
		Boolean bool
		Number  float32
		Number2 float64
		Int     int
		Int2    int8
		Int3    int16
		Int4    int32
		Int5    int64
		Int6    uint
		Int7    uint8
		Int8    uint16
		Int9    uint32
		Int10   uint64
		Struct  struct{}
		Slice   []interface{}
	)

	st := reflect.TypeOf(String)
	pt, err := getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, StringPropType, pt)

	st = reflect.TypeOf(Boolean)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, BoolPropType, pt)

	st = reflect.TypeOf(Number)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, NumberPropType, pt)

	st = reflect.TypeOf(Number2)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, NumberPropType, pt)

	st = reflect.TypeOf(Int)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int2)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int3)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int4)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int5)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int6)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int7)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int8)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int9)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Int10)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, IntPropType, pt)

	st = reflect.TypeOf(Struct)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, ObjectPropType, pt)

	st = reflect.TypeOf(Slice)
	pt, err = getPropTypeFromKind(st.Kind())
	assert.NoError(t, err)
	assert.Equal(t, ArrayPropType, pt)
}

func Test_GetPropsFromType(t *testing.T) {
	var (
		String string
		Array  []string
		Struct struct{}
	)

	tp := reflect.TypeOf(String)
	prop, err := getPropsFromType(tp)
	assert.NoError(t, err)
	assert.Equal(t, StringPropType, prop.Type)

	tp = reflect.TypeOf(Array)
	prop, err = getPropsFromType(tp)
	assert.NoError(t, err)
	assert.Equal(t, ArrayPropType, prop.Type)

	tp = reflect.TypeOf(Struct)
	prop, err = getPropsFromType(tp)
	assert.NoError(t, err)
	assert.Equal(t, ObjectPropType, prop.Type)

	type MyStruct struct {
		TestField       string                 `description:"test field description"`
		TestFieldSlice  []string               `description:"test string slice field description"`
		TestFieldStruct struct{ Inner string } `description:"test struct field description"`
	}
	tp = reflect.TypeOf(MyStruct{})
	prop, err = getPropsFromType(tp)
	assert.NoError(t, err)
	assert.Equal(t, ObjectPropType, prop.Type)

	testFieldProp, ok := prop.Properties["testField"]
	assert.True(t, ok)
	assert.Equal(t, "test field description", testFieldProp.Description)

	testFieldProp, ok = prop.Properties["testFieldSlice"]
	assert.True(t, ok)
	assert.Equal(t, "test string slice field description", testFieldProp.Description)

	testFieldProp, ok = prop.Properties["testFieldStruct"]
	assert.True(t, ok)
	assert.Equal(t, "test struct field description", testFieldProp.Description)
	innerPropFIeld, ok := testFieldProp.Properties["inner"]
	assert.True(t, ok)
	assert.Equal(t, StringPropType, innerPropFIeld.Type)

}
