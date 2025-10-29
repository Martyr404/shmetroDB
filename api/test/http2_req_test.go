package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

var url = "https://localhost:8080"

func MakeH2Req() {
	request2()
}

func request2() {
	path, _ := os.Getwd()
	client := &http.Client{
		Timeout: 100 * time.Millisecond,
	}
	//读取证书文件，正式环境无读取证书文件，因为本地测试是无法认证证书
	caCert, err := ioutil.ReadFile(path + "\\http_connect\\certificate.pem")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	//tls协议配置，InsecureSkipVerify认证证书是否跳过
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		//设置安全跳跃认证
		InsecureSkipVerify: true,
	}
	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Failed get: err:%s \n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed reading response body: %s\n", err)
		return
	}
	fmt.Printf("Get request2 response %d: %s %s\n", resp.StatusCode, resp.Proto, string(body))
}

// http1请求
func request1() {
	//http1
	_, err := http.Get(url)
	fmt.Println(err)
}

// http11请求
func request11() {
	path, _ := os.Getwd()
	client := &http.Client{}
	//证书文件
	caCert, err := ioutil.ReadFile(path + "\\certificate.pem")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		//设置安全跳跃认证
		InsecureSkipVerify: true,
	}
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
		return
	}
	fmt.Printf("Got response %d: %s %s\n", resp.StatusCode, resp.Proto, string(body))
}
