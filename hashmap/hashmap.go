package hashmap

import "strconv"

func NewHashMaps(keyName string, size int) map[string]int {
	m := make(map[string]int, size)
	for i := 0; i < size; i++ {
		s := strconv.Itoa(i)
		m[keyName+s] = i
	}
	return m
}
