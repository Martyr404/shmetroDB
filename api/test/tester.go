package main

import (
	"fmt"
	"shmetroDB/orm"
)

// func main() {
// 	body_data := map[string]string{"carriage_number": "052001"}
// 	byte_body, _ := json.Marshal(body_data)
// 	client := http.Client{}
// 	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:9987/api/calculate", bytes.NewReader(byte_body))
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, _ := io.ReadAll(resp.Body)

// 	fmt.Println(string(respBody))
// }

func main() {
	res, err := orm.ParseCarriageNumber("141182")
	fmt.Println(res, err)
}
