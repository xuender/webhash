package webhash

// Batch 批量摘要
func Batch(webURLs []string) []*Hash {
	ret := make([]*Hash, len(webURLs))
	for i, url := range webURLs {
		ret[i] = New(url)
	}
	return ret
}
