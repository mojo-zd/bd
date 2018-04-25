package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mojo-zd/bd/translate"
	"github.com/urfave/cli"
)

type TranslateResult struct {
	From         string
	To           string
	TransResults []TransResult `json:"trans_result"`
}

type TransResult struct {
	Src string
	Dst string
}

func translator(c *cli.Context) {
	result := &TranslateResult{}
	resp, err := http.Get(urlJoint(strings.Join(c.Args(), "%20"), "12345678"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, result)
	for _, r := range result.TransResults {
		fmt.Println(r.Dst)
	}
}

func urlJoint(query, random string) string {
	var from, to string
	if translate.IsChinese(query) {
		from = translate.ZH
		to = translate.En
	} else {
		from = translate.En
		to = translate.ZH
	}

	return fmt.Sprintf("%s?q=%s&from=%s&to=%s&appid=%s&salt=%s&sign=%s", translate.TranslatorURL, query,
		from, to, translate.AppId, random, translate.EncryptParams(random, query))
}
