package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	_url "net/url"
	"xjtlu-dorm-net-auth-helper/conf"
)

func Do(url, method string, params interface{}, result interface{}) error {
	url = conf.Get().URL + "/api/portal/v1" + url

	reqBody, err := json.Marshal(params)
	if err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to marshal JSON:", err, params)
		return errors.New("failed to marshal JSON")
	}
	fmt.Println("[DEBUG/REQUEST] JSON marshalled.")

	buffer := bytes.NewBuffer(reqBody)

	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to create a new HTTP request:", err, method, url, params)
		return errors.New("faild to create a new HTTP request")
	}
	fmt.Println("[DEBUG/REQUEST] New HTTP request created.")

	if err := SetHeader(request); err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to set header:", err)
		return errors.New("failed to set header")
	}
	fmt.Println("[DEBUG/REQUEST] Request header is set.")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to send HTTP request:", err, method, url, params)
		return errors.New("failed to send HTTP request")
	}
	defer response.Body.Close()
	fmt.Println("[DEBUG/REQUEST] HTTP request sended.")

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to read HTTP response:", err, method, url, params)
		return errors.New("failed to read HTTP response")
	}
	fmt.Println("[DEBUG/REQUEST] HTTP response readed.")

	err = json.Unmarshal(resBody, result)
	if err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to unmarshal JSON:", err, method, url, params)
		return errors.New("failed to unmarshal JSON")
	}
	fmt.Println("[DEBUG/REQUEST] JSON Unmarshalled.")

	return nil
}

func SetHeader(request *http.Request) error {
	headers := make(map[string]string)
	headers["Accept"] = "application/json, text/javascript, */*; q=0.01"
	headers["Accept-Encoding"] = "gzip, deflate"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"
	headers["Cache-Control"] = "no-cache"
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8" // ???

	u, err := _url.Parse(conf.Get().URL)
	if err != nil {
		fmt.Println("[ERROR/REQUEST] Failed to parse URL:", err, conf.Get().URL)
		return errors.New("failed to parse URL")
	}
	fmt.Println("[DEBUG/REQUEST] URL parsed.")

	headers["Host"] = u.Host

	headers["Origin"] = conf.Get().URL
	headers["Pragma"] = "no-cache"
	headers["Proxy-Connection"] = "keep-alive"
	headers["Referer"] = conf.Get().URL + "/portal/"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0"

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return nil
}
