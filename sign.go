package superpay

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strings"
)

func sign(paramsKv map[string]string, key string) string {
	u := url.Values{}
	for k, v := range paramsKv {
		if v == "" {
			continue
		}
		u.Set(k, v)
	}
	s := u.Encode()
	s += "&key=" + key
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
