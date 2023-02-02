package fnReflect

import (
	"fmt"
	"reflect"
)

// IsStruct
// v가 struct 타입인지 확인한다.
func IsStruct(v any) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Pointer:
		return reflect.TypeOf(v).Elem().Kind() == reflect.Struct
	case reflect.Struct:
		return true
	default:
		return false
	}
}

// GetStructNm
// v가 struct 일 경우 그 이름을 반환한다.
// v가 struct 가 아닐경우 빈값 "" 을 반환한다.
func GetStructNm(v any) (string, error) {
	if !IsStruct(v) {
		return "", fmt.Errorf("v is not struct")
	}
	return reflect.TypeOf(v).Name(), nil
}

func Kind[T any]() reflect.Kind {
	return reflect.TypeOf(new(T)).Kind()
}
