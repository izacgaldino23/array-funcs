package arrayfuncs

import (
	"reflect"
	"strconv"
)

// AnyToString parse almost all values to string
func AnyToString(value any) (converted string) {
	if value == nil {
		return
	}

	valueOf := reflect.ValueOf(value)
	typeOf := reflect.TypeOf(value)

	if typeOf.Kind() == reflect.Pointer {
		valueOf = reflect.ValueOf(value).Elem()
		typeOf = reflect.TypeOf(value).Elem()
	}

	switch typeOf.Kind() {
	case reflect.Int:
		converted = strconv.FormatInt(int64(valueOf.Interface().(int)), 10)
	case reflect.Bool:
		converted = strconv.FormatBool(valueOf.Interface().(bool))
	case reflect.Float32:
		converted = strconv.FormatFloat(float64(valueOf.Interface().(float32)), 'f', -1, 32)
	case reflect.Float64:
		converted = strconv.FormatFloat(valueOf.Interface().(float64), 'f', -1, 64)
	case reflect.String:
		converted = valueOf.Interface().(string)
	case reflect.Struct:
		temp := reflect.ValueOf(value)
		toString := temp.MethodByName("ToString")

		if toString.IsValid() {
			result := toString.Call([]reflect.Value{})
			converted = result[0].Interface().(string)
		}
	}

	return
}
