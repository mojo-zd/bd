package translate

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

var (
	AppId         = "20180425000150381"
	Secret        = "Kb4WE5T8Lq3r6PbbaIEh"
	TranslatorURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	ZH            = "zh"
	En            = "en"
)

// EncryptParams ...
func EncryptParams(random, query string) string {
	s := AppId + query + random + Secret
	h := md5.New()
	io.WriteString(h, s)
	m5 := fmt.Sprintf("%x", h.Sum(nil))
	return m5
}

// IsChinese ...
func IsChinese(value string) bool {
	reg := regexp.MustCompile("^[\u4e00-\u9fa5]$")
	for _, v := range value {
		if reg.MatchString(string(v)) {
			return true
		}
	}
	return false
}
