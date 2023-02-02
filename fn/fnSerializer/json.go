package fnSerializer

import (
	"encoding/json"
)

func ToJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}

func ToStruct[T any](v []byte) (*T, error) {
	res := new(T)
	if err := json.Unmarshal(v, res); err != nil {
		return nil, err
	}
	return res, nil
}

//
//func ToParams(v any) (res string, err error) {
//
//	vo := reflect.ValueOf(v)
//
//	switch vo.Kind() {
//	case reflect.Pointer, reflect.UnsafePointer:
//		if vo.IsNil() {
//			return "", fmt.Errorf("v is nil")
//		}
//		return ToParams(vo.Elem().Interface())
//	case reflect.String:
//		return fmt.Sprintf("%s", v), nil
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
//		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//		return fmt.Sprintf("%d", v), nil
//	case reflect.Float32, reflect.Float64:
//		return fmt.Sprintf("%f", v), nil
//	case reflect.Slice, reflect.Array:
//		if vo.Len() == 0 {
//			return "[]", nil
//		}
//
//		res += "["
//		for i := 0; i < vo.Len(); i++ {
//			elem := vo.Index(i)
//			var s string
//			if s, err = ToParams(elem); err != nil {
//				return "", err
//			}
//			res += fmt.Sprintf("%s,", s)
//		}
//
//		res += "],"
//	}
//
//	return
//}
