package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	_url "net/url"
	"xjtlu-dorm-net-auth-helper/conf"
	"xjtlu-dorm-net-auth-helper/logger"
)

func Do(url, method string, params interface{}, result interface{}) error {
	url = conf.Get().URL + "/api/portal/v1" + url

	logger.Debug("Marshalling JSON")
	reqBody, err := json.Marshal(params)
	if err != nil {
		logger.Debug("Failed to marshal JSON: %s", err)
		return err
	}

	buffer := bytes.NewBuffer(reqBody)

	logger.Debug("Creating HTTP request: %s, %s", method, url)
	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		logger.Debug("Failed to create a new HTTP request: %s", err)
		return err
	}

	logger.Debug("Setting request header")
	if err := SetHeader(request); err != nil {
		logger.Debug("Failed to set header: %s", err)
		return err
	}

	client := &http.Client{}
	logger.Debug("Sending HTTP request")
	response, err := client.Do(request)
	if err != nil {
		logger.Debug("Failed to send HTTP request: %s", err)
		return err
	}
	defer response.Body.Close()

	logger.Debug("Reading HTTP response")
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Debug("Failed to read HTTP response: %s", err)
		return err
	}

	logger.Debug("Unmarshalling JSON")
	err = json.Unmarshal(resBody, result)
	if err != nil {
		logger.Debug("Failed to unmarshal JSON: %s", err)
		return err
	}

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
		logger.Error("Failed to parse URL: %s", err)
		return err
	}

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
