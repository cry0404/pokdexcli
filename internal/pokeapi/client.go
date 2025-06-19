package pokeapi


import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2/location-area"

// 这里是如何构建一个最基础的客户端，老是记不住

type Client struct {
	httpClient *http.Client
}

//封装创建逻辑

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) GetLocationAreas(url string)(*LocationResponse, error) {
	if url == ""{
		
		url = baseURL 
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body failed: %w", err)
	}
	var locationResp LocationResponse
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &locationResp, nil
	//整体的任务就是返回一个结构体数组
}