package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	_url "net/url"
	"strings"
	"xjtlu-dorm-net-helper/conf"

	"github.com/gorilla/schema"
)

func Do(url, method string, body interface{}, result interface{}) error {
	url = conf.Get().URL + "/api/portal/v1" + url

	encoder := schema.NewEncoder()
	values := _url.Values{}

	if err := encoder.Encode(body, values); err != nil {
		return err
	}

	reader := strings.NewReader(values.Encode())

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
	}

	if err := SetHeader(request); err != nil {
		return err
	}

	fmt.Println("[HTTP] POST ", url)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return err
	}

	return nil
}

func SetHeader(request *http.Request) error {
	headers := make(http.Header)
	headers.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	headers.Set("Accept-Encoding", "gzip, deflate")
	headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	headers.Set("Cache-Control", "no-cache")
	headers.Set("Content-Type", "applicationapplication/x-www-form-urlencoded; charset=UTF-8")

	u, err := _url.Parse(conf.Get().URL)
	if err != nil {
		return err
	}
	headers.Set("Host", u.Host)

	headers.Set("Origin", conf.Get().URL)
	headers.Set("Pragma", "no-cache")
	headers.Set("Proxy-Connection", "keep-alive")
	headers.Set("Referer", conf.Get().URL+"/portal/")
	headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0")

	for key, values := range headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	return nil
}
