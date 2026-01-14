package arrays

func Reduce[T, U any](collection []T, fn func(U, T) U, initialValue U) U {
	result := initialValue
	for _, elem := range collection {
		result = fn(result, elem)
	}
	return result
}

func Find[T any](items []T, predicate func(T) bool) (value T, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
