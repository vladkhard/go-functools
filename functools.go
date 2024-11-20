package functools

import (
	"fmt"
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

func (iterator Iterator[int, V]) FilterSlice(filter func(index int, value V) bool) Iterator[int, V] {
	return func(yield func(int, V) bool) {
		internalIndex := 0
		var internalIndexBox int = any(internalIndex).(int)
		for index, value := range iterator {
			if filter(index, value) {
				yield(internalIndexBox, value)
				internalIndex++
				internalIndexBox = any(internalIndex).(int)
			}
		}
	}
}

func (iterator Iterator[K, V]) FilterMap(filter func(key K, value V) bool) Iterator[K, V] {
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

func (iterator Iterator[K, V]) String() string {
	stringBuffer := ""
	for key, value := range iterator {
		if stringBuffer == "" {
			stringBuffer = fmt.Sprintf("%v: %v", key, value)
		} else {
			stringBuffer += fmt.Sprintf(", %v: %v", key, value)
		}
	}
	return fmt.Sprintf("Iterator[%s]", stringBuffer)
}
