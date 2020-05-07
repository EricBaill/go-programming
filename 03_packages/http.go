package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//httpGet()
	//httpPost()
	httpClient()
}

func httpGet() {

	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("body:", string(body))
}

func httpPost() {

	//resp, err := http.Post("http://www.baidu.com",
	//	"application/json",
	//	strings.NewReader("http://github.com/meetbetter/go-programming"))
	reqBody := "{\"money\":\"many\"}"
	resp, err := http.Post("http://www.baidu.com", "application/json",
		bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println(string(respBody))
}

func httpClient() {
	v := url.Values{}
	v.Set("level", "p8")
	v.Set("money", "many")
	body := strings.NewReader(v.Encode()) //form编码
	client := &http.Client{}              //client
	reqest, err := http.NewRequest("POST", "http://www.baidu.com", body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//给一个key设定为响应的value.
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") //必须设定该参数,POST参数才能正常提交

	resp, err := client.Do(reqest) //发送请求
	if err != nil {
		fmt.Println("client.Do error ", err.Error())
		return

	}
	defer resp.Body.Close() //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}
