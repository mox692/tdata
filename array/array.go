package array

import "strconv"

func NewArrayInt(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = i
	}
	return a
}

func NewArrayString(size int) []string {
	a := make([]string, size)
	for i := 0; i < size; i++ {
		s := strconv.Itoa(i)
		a[i] = s
	}
	return a
}

func NewArrayInterface(size int, mapFunc func(in int) interface{}) []interface{} {
	a := make([]interface{}, size)
	for i := 0; i < size; i++ {
		a[i] = mapFunc(i)
	}
	return a
}
