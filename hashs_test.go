package webhash

import "testing"

func TestHashs_Add(t *testing.T) {
	hs := Hashs{}
	h := &Hash{ID: 1}
	hs.Add(h)
	if len(hs) != 1 {
		t.Errorf("add 错误:%d", len(hs))
	}
}
