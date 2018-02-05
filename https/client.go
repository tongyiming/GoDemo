package main

import (
	"net/http"
	"crypto/tls"
	"fmt"
	"io/ioutil"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//设置InsecureSkipVerify为true使客户端不对服务端进行校验
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8080")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
