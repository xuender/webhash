package webhash

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mitchellh/mapstructure"
)

// Hash 网页摘要
type Hash struct {
	ID    uint64
	URL   string
	Error error `yaml:"-"`
	Sum   uint64
	Time  int64
}

func (h *Hash) String() string {
	if h.Error == nil {
		return fmt.Sprintf("摘要: %d, 网址: [%s] 时间: %s", h.Sum, h.URL, time.Unix(h.Time, 0))
	}
	return fmt.Sprintf("网址: [%s] 错误: %s", h.URL, h.Error)
}

// Get 请求
func (h *Hash) Get() (keep bool) {
	resp, err := http.Get(h.URL)
	if err != nil {
		h.Error = err
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		h.Error = fmt.Errorf("网络异常:%d", resp.StatusCode)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.Error = err
		return
	}
	sum := Siphash(bytes)
	keep = h.Sum == sum
	h.Sum = sum
	return
}

// New 新建摘要
func New(webURL string) (h *Hash) {
	h = &Hash{
		ID:   Siphash([]byte(webURL)),
		URL:  webURL,
		Time: time.Now().Unix(),
	}
	if !govalidator.IsURL(webURL) {
		h.Error = errors.New("错误的URL")
		return
	}
	h.Get()
	return
}

// Parse 解析
func Parse(m interface{}) (*Hash, error) {
	var hash Hash
	err := mapstructure.Decode(m, &hash)
	return &hash, err
}
