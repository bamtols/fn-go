package fnParams

func Pick[T any](v []T) T {
	if len(v) == 0 {
		return *new(T)
	} else {
		return v[0]
	}
}

func PickWithDefault[T any](v []T, def T) T {
	if len(v) == 0 {
		return def
	} else {
		return v[0]
	}
}
