package hashmap

import "testing"

func Benchmark_NewHashMaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewHashMaps("key", 100000)
	}
}
