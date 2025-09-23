package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 定义请求体结构（与服务器端对应）
type DemoRequestBody struct {
	Carriage_number string `json:"carriage_number"` // 必须带JSON标签，否则序列化后字段名可能错误
}

func RaiseDemoRequest() {
	client := http.Client{}

	// 1. 构建请求体数据（JSON格式）
	reqBody := DemoRequestBody{
		Carriage_number: "080071", // 要发送的编号
	}

	// 2. 序列化为JSON字符串（关键：确保发送的是JSON）
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("JSON序列化失败: %s\n", err.Error())
		return
	}

	// 3. 创建POST请求（改用POST方法，符合HTTP规范）
	req, err := http.NewRequest("POST", "http://127.0.0.1:9987/calculate", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("创建请求失败: %s\n", err.Error())
		return
	}

	// 4. 设置正确的Content-Type（告诉服务器是JSON格式）
	req.Header.Set("Content-Type", "application/json")

	// 5. 发送请求并处理响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %s\n", err.Error())
		return
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 6. 打印响应头
	fmt.Println("响应头:")
	for key, value := range resp.Header {
		fmt.Printf("%s: %v\n", key, value)
	}

	// 7. 打印响应体
	bodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应体失败: %s\n", err.Error())
		return
	}
	fmt.Println("\n响应体:")
	fmt.Println(string(bodyRaw))
}

func main() {
	RaiseDemoRequest()
}
