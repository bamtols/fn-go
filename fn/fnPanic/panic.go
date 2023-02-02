package fnPanic

import "log"

func HasError(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func HasErrorOrValue[T any](v T, e error) T {
	if e != nil {
		log.Panic(e)
	}
	return v
}
