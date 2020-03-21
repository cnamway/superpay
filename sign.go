package superpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

func sign(params map[string]string, key string) string {
	s := mapToUrl(params, key)
	fmt.Println(s)
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func mapToUrl(params map[string]string, key string) string {
	u := url.Values{}
	for k, v := range params {
		if v == "" {
			continue
		}
		u.Set(k, v)
	}
	s := u.Encode()
	s += "&key=" + key
	return s
}
