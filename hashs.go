package webhash

// Hashs 摘要列表
type Hashs []*Hash

// Add 增加
func (h *Hashs) Add(hash *Hash) {
	for _, o := range *h {
		if o.ID == hash.ID {
			return
		}
	}
	*h = append(*h, hash)
}

// NewHashs 新建列表
func NewHashs(i interface{}) Hashs {
	hashs := Hashs{}
	for _, h := range i.([]interface{}) {
		if hash, err := Parse(h); err == nil {
			hashs.Add(hash)
		}
	}
	return hashs
}
