package fnReflect

// GetValueFromPointer
// v 는 포인터, 또는 값
// v 가 포인터값이면 nil 이 아니면 값을 꺼내고, nil 이면 panic
// v 가 value 라면 value 반환.
//func GetValueFromPointer[T any](v any) T {
//	switch reflect.TypeOf(v).Kind() {
//	case reflect.Pointer:
//		v := reflect.ValueOf(v)
//		if v.IsNil() {
//			panic(fmt.Errorf("v is nil"))
//		}
//
//		res, isOk := v.Elem().Interface().(T)
//		if !isOk {
//			panic(fmt.Errorf("v is not T"))
//		}
//
//		return res
//	default:
//		_, isOk := v.(T)
//		if isOk {
//			return v
//		} else {
//			panic(fmt.Errorf("v is not T"))
//		}
//	}
//}

func ToPointer[T any](v T) *T {
	return &v
}
