package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	body_data := map[string]string{"carriage_number": "02011"}
	byte_body, _ := json.Marshal(body_data)
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:9987/api/calculate", bytes.NewReader(byte_body))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respBody))
}
