package webhash

import "github.com/dchest/siphash"

var _KEY = []byte("xuender,20200509")

// Siphash 摘要
func Siphash(bytes []byte) uint64 {
	hash := siphash.New(_KEY)
	hash.Write(bytes)
	return hash.Sum64()
}
