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
	sync.RWMutex
	ID        int
	LinksShort      map[string]string
	LinksOriginal      map[string]string
	ShemeHost string
}

// Shorten makes a shorе url from an orirignal url
func (u *Urls) Shorten(orirignal string) string {
	u.Lock()
	defer u.Unlock()
	if val, ok := u.LinksOriginal[orirignal]; ok {
		return val
	}
	u.ID++
	shortURL := fmt.Sprintf("%s/%s", u.ShemeHost, u.Encode(u.ID))
	u.LinksOriginal[orirignal] = shortURL
	u.LinksShort[shortURL] = orirignal
	return shortURL
}

// Resolver makes an orirignal url from a short url
func (u *Urls) Resolver(short string) string {
	u.RLock()
	defer u.RUnlock()
	if val, ok := u.LinksShort[short]; ok {
		return val
	}
	return ""
}

// Encode int to base62
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
