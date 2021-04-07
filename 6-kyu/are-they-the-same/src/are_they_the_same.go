package src

import (
	"reflect"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil && array2 == nil {
		return false
	}

	for i, v := range array1 {
		array1[i] *= v
	}

	sort.Ints(array1)
	sort.Ints(array2)
	return reflect.DeepEqual(array1, array2)
}
