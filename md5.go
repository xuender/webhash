package webhash

import "crypto/md5"

// Md5 摘要
func Md5(bytes []byte) []byte {
	hash := md5.New()
	hash.Write(bytes)
	return hash.Sum(nil)
}
