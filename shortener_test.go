package shortener

import (
	"strings"
	"testing"
)

func TestShort(t *testing.T) {

	urls := urls{
		ID:        20000,
		urls:      make(map[string]string),
		shemeHost: "https://otus.ru",
	}

	orirginalURL := "https://ya.ru/hello/word/"
	shortURL := urls.Shorten(orirginalURL)
	resolveURL := urls.Resolver(shortURL)
	if strings.Compare(orirginalURL, resolveURL) != 0 {
		t.Errorf("%s != %s", orirginalURL, resolveURL)
	}

	resolveURL = urls.Resolver("http://fee.com/test")
	if strings.Compare("", resolveURL) != 0 {
		t.Error("resolveURL must be ''")
	}

}
