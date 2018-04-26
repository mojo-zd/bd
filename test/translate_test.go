package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/astaxie/beego/utils"
	"github.com/mojo-zd/bd/translate"
)

var (
	appid  = "20180425000150381"
	secret = "Kb4WE5T8Lq3r6PbbaIEh"
	salt   = "12345678"
	u      = "\u4e2d\u56fd\u0020\u82af\u7247"
)

func Test(t *testing.T) {
	s := appid + u + salt + secret
	h := md5.New()
	fmt.Println("===>", s)
	io.WriteString(h, s)
	fmt.Printf("%x", h.Sum(nil))
}

func TestT(t *testing.T) {
	var query string = "人"
	var random string = "12345678"
	var from, to string
	if translate.IsChinese(query) {
		from = translate.ZH
		to = translate.En
	} else {
		from = translate.En
		to = translate.ZH
	}
	url := fmt.Sprintf("%s?q=%s&from=%s&to=%s&appid=%s&salt=%s&sign=%s", translate.TranslatorURL, query,
		from, to, translate.AppId, random, translate.EncryptParams(random, query))
	fmt.Println(url)
	resp, _ := http.Post(url, "application/json", nil)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	utils.Display("", string(body))
}

func TestYD(t *testing.T) {
	var query string = "人"
	var random string = "12345678"
	var from, to string
	if translate.IsChinese(query) {
		from = "zh_CHS"
		to = "EN"
	} else {
		from = "EN"
		to = "zh_CHS"
	}
	url := fmt.Sprintf("%s?q=%s&from=%s&to=%s&appKey=%s&salt=%s&sign=%s", translate.YdTranslatorURL, query,
		from, to, translate.YdAppId, random, encryptParams(random, query))
	fmt.Println(url)
	resp, _ := http.Post(url, "application/json", nil)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	utils.Display("", string(body))
}

func encryptParams(random, query string) string {
	s := translate.YdAppId + query + random + translate.YdSecret
	h := md5.New()
	io.WriteString(h, s)
	m5 := fmt.Sprintf("%x", h.Sum(nil))
	return m5
}
