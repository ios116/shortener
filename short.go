package shortener

import (
	"fmt"
	"sync"
)

func init() {
	// url := url {
	// 	ID : 0,
	// 	Urls: make(map[string]string),
	// }
}

// Shortener interface
type Shortener interface {
	Shorten(orirignal string) string
	Resolver(short string) string
}

// Urls ID - start value, Links is map short an original links
type Urls struct {
	sync.Mutex
	ID        int
	Links      map[string]string
	ShemeHost string
}

// Shorten makes a shorе url from an orirignal url
func (u *Urls) Shorten(orirignal string) string {
	u.Lock()
	defer u.Unlock()
	if val, ok := u.Links[orirignal]; ok {
		return val
	}
	u.ID++
	shortURL := fmt.Sprintf("%s/%s", u.ShemeHost, u.Encode(u.ID))
	u.Links[orirignal] = shortURL
	return shortURL
}

// Resolver makes an orirignal url from a short url
func (u *Urls) Resolver(short string) string {

	u.Lock()
	defer u.Unlock()
	for k, v := range u.Links {
		if v == short {
			return k
		}
	}
	return ""
}

// Encode to int to base62
func (u *Urls) Encode(id int) string {
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
