package reflection

import (
	"reflect"
)

func Walk(x any, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := range val.NumField() {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			v, ok := val.Recv()
			if !ok {
				break
			}

			walkValue(v)
		}
	case reflect.Func:
		// Only call zero-argument functions — others may have side effects
		if val.Type().NumIn() != 0 {
			return
		}
		for _, v := range val.Call(nil) {
			walkValue(v)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
