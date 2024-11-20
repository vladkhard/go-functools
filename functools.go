package functools

import (
	"iter"
	"maps"
	"slices"
)

type Iterator[K, V any] iter.Seq2[K, V]

func SliceIterator[K any](from []K) Iterator[int, K] {
	return Iterator[int, K](slices.All(from))
}

func MapIterator[K comparable, V any](from map[K]V) Iterator[K, V] {
	return Iterator[K, V](maps.All(from))
}

func (iterator Iterator[K, V]) Filter(filter func(key K, value V) bool) Iterator[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range iterator {
			if filter(key, value) {
				yield(key, value)
			}
		}
	}
}

func (iterator Iterator[K, V]) Map(function func(key K, value V) (K, V)) Iterator[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range iterator {
			key, value = function(key, value)
			yield(key, value)
		}
	}
}
