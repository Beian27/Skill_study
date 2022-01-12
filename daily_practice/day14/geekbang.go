package main

import (
	"fmt"
	"github.com/onyas/go-browsercookie"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var cookieJar, _ = browsercookie.Chrome("youlexue.yonyou.com")
	var client = &http.Client{Jar: cookieJar}

	request, _ := http.NewRequest("POST", "http://youlexue.yonyou.com/youlexuehome.htm", strings.NewReader("{}"))

	request.Header.Set("Origin", "http://youlexue.yonyou.com")
	request.Header.Set("Referer", "http://youlexue.yonyou.com/kng/kngtags/knowledgetagsmgmt.htm")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
	request.Header.Set("X-Real-IP", "120.53.129.107")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Content-Length", "22")

	res, _ := client.Do(request)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}