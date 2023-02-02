package fnArray

import "fmt"

func Find[T any](list []T, fn func(v T) bool) (T, error) {
	for _, t := range list {
		if fn(t) {
			return t, nil
		}
	}
	return *new(T), fmt.Errorf("notFoundData")
}
