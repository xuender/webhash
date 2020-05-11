package webhash

import (
	"fmt"
	"testing"
)

func ExampleHighwayhash() {
	fmt.Printf("nil=%d\n", Highwayhash(nil))
	fmt.Printf("test=%d\n", Highwayhash([]byte("test")))

	// Output:
	// nil=2319126219498428999
	// test=16819604689299744157
}

var testBytes = make([]byte, 50000)

// func BenchmarkSiphash(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Siphash(testBytes)
// 	}
// }
func BenchmarkHighwayhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Highwayhash(testBytes)
	}
}
