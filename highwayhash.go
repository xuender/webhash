package webhash

import (
	"github.com/minio/highwayhash"
)

var _Key32 = []byte("Webhash:author xuender, 20200511")

// Highwayhash 摘要
func Highwayhash(bytes []byte) uint64 {
	hash, _ := highwayhash.New64(_Key32)
	hash.Write(bytes)
	return hash.Sum64()
}
