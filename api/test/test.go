package test

import (
	"fmt"
	"io"
	"net/http"
)

type DemoRequestBody struct {
	Carriage_number string
}

func (b DemoRequestBody) Read(content []byte) (int, error) {
	n := copy(content, []byte(b.Carriage_number))
	if n < len(b.Carriage_number) {
		return n, nil
	}
	//超出容量
	return n, io.EOF
}

func raiseDemoRequest() {
	client := http.Client{}
	body := DemoRequestBody{}
	req, err := http.NewRequest("GET", "127.0.0.1:9987", body)
	if err != nil {
		fmt.Printf("创建请求失败:%s\n", err.Error())
		return
	}
	req.Header.Set("Content-Type", "text/plain")
	resp, _ := client.Do(req)
}
