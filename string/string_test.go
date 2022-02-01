package string

import (
	"fmt"
	"testing"
)

func TestNewStringSlice(t *testing.T) {
	size := 30
	strs := NewAsciiStringSlice(size)
	for i := 0; i < size; i++ {
		fmt.Println(strs[i])
	}
}

func BenchmarkNewStringSlice_1000(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewAsciiStringSlice(100000)
	}
	b.StopTimer()
}
