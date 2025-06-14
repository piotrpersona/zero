package zero

func Map[T, R any](coll []T, mapper func(T) R) []R {
	ret := make([]R, 0, len(coll))
	for _, elem := range coll {
		ret = append(ret, mapper(elem))
	}
	return ret
}

func MapWithErr[T, R any](coll []T, mapper func(T) (R, error)) ([]R, error) {
	ret := make([]R, 0, len(coll))
	for _, elem := range coll {
		mapped, err := mapper(elem)
		if err != nil {
			return nil, err
		}
		ret = append(ret, mapped)
	}
	return ret, nil
}

func Filter[T any](coll []T, pred func(T) bool) []T {
	ret := make([]T, 0)
	for _, elem := range coll {
		if pred(elem) {
			ret = append(ret, elem)
		}
	}
	return ret
}
