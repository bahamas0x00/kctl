package common

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var (
	once   sync.Once
	client *http.Client
)

// return a http client
func getClient() *http.Client {
	once.Do(func() {
		client = &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 不验证证书
				},
			},
		}
	})
	return client
}

// common request sender
func SendRequest(method, apiEndpoint string, pathComponents []string, data interface{}) (*http.Response, error) {
	httpClient := getClient()
	// 构造 URL 路径
	urlPath, err := url.JoinPath(apiEndpoint, pathComponents...)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %v", err)
	}

	var body io.Reader
	if data != nil {
		// 如果有数据，则需要将其编码为 JSON
		bodyData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("invalid data: %v", err)
		}
		body = bytes.NewBuffer(bodyData)
	}

	fmt.Printf("Method: %s --> Request Path: %s\n", method, urlPath)

	// 创建 HTTP 请求
	httpRequest, err := http.NewRequest(method, urlPath, body)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	// 发送请求并返回响应
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}

func IsStringSet(s string) bool {
	return s != ""
}

// save response content to file
func SaveResponseToFile(response *http.Response, outputFile string) error {
	// create or open file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// write response content to file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	fmt.Println("写入文件成功")
	return nil
}
