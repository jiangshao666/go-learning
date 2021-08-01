package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 发起http请求

// 基础get，不带参数
func getNoArg() {
	resp, err := http.Get("https://www.bilibili.com/")
	if err != nil {
		fmt.Printf("http Get failed, err:%v", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http Get ReadAll failed, err:%v", err)
		return
	}
	fmt.Println(string(b))
}

// 带参数get请求
func getWithArg() {
	apiURL := "http://127.0.0.1:8090/products"
	data := url.Values{}
	data.Set("productName", "华为荣耀P40")
	data.Set("pageNum", "20")
	queryData := data.Encode()
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		fmt.Printf("ParseRequestURI failed, err:%v", err)
		return
	}
	u.RawQuery = queryData

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("http Get failed, err:%v", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http Get ReadAll failed, err:%v", err)
		return
	}
	fmt.Println(string(b))

}

// post请求
func postWithArg() {
	apiURL := "http://127.0.0.1:8090/login"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"

	contentType := "application/json"
	data := `{"user":"张三","pass":123456}`
	resp, err := http.Post(apiURL, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf(" post data failed, err:%v", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http Get ReadAll failed, err:%v", err)
		return
	}
	fmt.Println(string(b))

}

func main() {
	//getNoArg()
	//getWithArg()

	postWithArg()
}
