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

type HttpResponse struct {
	StatusCode int
	Headers    map[string][]string
	Body       io.Reader
}

var once sync.Once
var client *http.Client

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
func SendRequest(method, apiEndpoint string, pathComponents []string, data interface{}) (*HttpResponse, error) {
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
	defer resp.Body.Close()

	return &HttpResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       resp.Body,
	}, nil
}

// ReadJsonFromFile read data fron json file and send request
func ReadJsonFromFileAndSendRequest(method, apiEndpoint string, pathComponents []string, filePath string) (*HttpResponse, error) {
	// 读取 JSON 文件
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	// 将文件内容解析成适当的格式，假设文件内容就是我们想要发送的 JSON 数据
	var jsonData interface{}
	if err := json.Unmarshal(fileContent, &jsonData); err != nil {
		return nil, fmt.Errorf("error unmarshaling file content: %v", err)
	}

	// 发送请求
	return SendRequest(method, apiEndpoint, pathComponents, jsonData)
}

func IsStringSet(s string) bool {
	return s != ""
}

// save response content to file
func SaveResponseToFile(response *HttpResponse, outputFile string) error {
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

	fmt.Println("HTTP 响应状态码:", response.StatusCode)
	return nil
}



// Convert a slice of any struct type to a slice of pointers to those structs
func ConvertToPointers[T any](data []T) []*T {
	pointers := make([]*T, len(data))
	for i := range data {
		pointers[i] = &data[i]
	}
	return pointers
}
