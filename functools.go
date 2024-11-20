package functools

import "iter"

func Filter[K, V any](iterable iter.Seq2[K, V], filter func(key K, value V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range iterable {
			if filter(key, value) {
				yield(key, value)
			}
		}
	}
}

func Map[K, V any](iterable iter.Seq2[K, V], function func(key K, value V) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range iterable {
			key, value = function(key, value)
			yield(key, value)
		}
	}
}

func ToSlice[V any](iterable iter.Seq2[int, V]) []V {
	var result []V
	for _, value := range iterable {
		result = append(result, value)
	}
	return result
}

func ToMap[K comparable, V any](iterable iter.Seq2[K, V]) map[K]V {
	result := make(map[K]V)
	for key, value := range iterable {
		result[key] = value
	}
	return result
}
