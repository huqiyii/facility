package conve

import "reflect"

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return true
	}
	return false
}
