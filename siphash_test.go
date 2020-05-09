package webhash

import "fmt"

func ExampleSiphash() {
	fmt.Printf("nil=%d\n", Siphash(nil))
	fmt.Printf("test=%d\n", Siphash([]byte("test")))

	// Output:
	// nil=15430654108413921344
	// test=5010450721137614759
}
