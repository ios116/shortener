package shortener

import (
	"fmt"
	"sync"
)

func init() {
	// url := url {
	// 	ID : 0,
	// 	urls: make(map[string]string),
	// }
}

// Shortener interface
type Shortener interface {
	Shorten(orirignal string) string
	Resolver(short string) string
}

type urls struct {
	sync.Mutex
	ID        int
	urls      map[string]string
	shemeHost string
}

func (u *urls) Shorten(orirignal string) string {
	u.Lock()
	defer u.Unlock()
	if val, ok := u.urls[orirignal]; ok {
		return val
	}
	u.ID++
	shortURL := fmt.Sprintf("%s/%s", u.shemeHost, u.Encode(u.ID))
	u.urls[orirignal] = shortURL
	return shortURL
}

func (u *urls) Resolver(short string) string {

	u.Lock()
	defer u.Unlock()
	for k, v := range u.urls {
		if v == short {
			return k
		}
	}
	return ""
}

func (u *urls) Encode(id int) string {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base := len(alphabet)
	b := make([]byte, 0)
	for id > 0 {
		r := id % base
		id /= base
		b = append([]byte{alphabet[r]}, b...)
	}
	return string(b)
}
