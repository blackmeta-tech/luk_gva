package utils

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const ()

var userAgent = []string{
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; it; rv:1.8.1.11) Gecko/20071127 Firefox/2.0.0.11",
	"Opera/9.25 (Windows NT 5.1; U; en)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Kubuntu)",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.12) Gecko/20070731 Ubuntu/dapper-security Firefox/1.5.0.12",
	"Lynx/2.8.5rel.1 libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/1.2.9",
	"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.04 Chromium/16.0.912.77 Chrome/16.0.912.77 Safari/535.7",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:10.0) Gecko/20100101 Firefox/10.0",
}

func HttpRequest(addr, method string, params url.Values) (resp *http.Response, err error) {
	Addr := addr + method
	Url, ok := url.Parse(Addr)
	if ok != nil {
		return
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println("=====", urlPath)
	if resp, err = http.Get(urlPath); err != nil {
		global.GVA_LOG.Error("Get请求出现错误", zap.Error(err))
		return
	}
	return
}

func HttpRequestPost(addr, method string, params map[string]interface{}) (resp *http.Response, err error) {
	Addr := addr + method
	bytes, _ := json.Marshal(params)
	stringData := string(bytes)
	payload := strings.NewReader(stringData)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, Addr, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func HttpRequestGetBsc(addr, method string, params url.Values) (resp *http.Response, err error) {
	Addr := addr + method
	Url, ok := url.Parse(Addr)
	if ok != nil {
		return
	}
	Url.RawQuery = params.Encode()
	urlPathWithParams := Url.String()
	fmt.Println(urlPathWithParams)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlPathWithParams, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var userA string
	randI := rand.Intn(13)
	if randI <= 12 {
		userA = userAgent[randI]
	} else {
		userA = userAgent[0]
	}
	fmt.Println(userA)
	req.Header.Add("User-Agent", userA)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func Unmarshal(body io.ReadCloser, tempData interface{}) {
	Body, _ := ioutil.ReadAll(body)
	json.Unmarshal(Body, &tempData)
}
