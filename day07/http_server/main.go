package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 不带参数GET
func index(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(`welcome to the edan! <a href="/login">login</a>`))
}

// 获取GET请求参数
func getProducts(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	fmt.Println("body", req.Body)
	fmt.Println("header", req.Header)
	fmt.Println("url", req.URL)
	fmt.Println("requrl", req.RequestURI)
	fmt.Println("uri", req.URL.Query())
	queryParam := req.URL.Query()
	pageNum := queryParam.Get("pageNum")
	productName := queryParam.Get("productName")
	fmt.Println("productName and pageNum", productName, pageNum)

	resp := `{"status":"ok"}`
	rw.Write([]byte(resp))
}

// 获取Post请求参数
func login(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	// 如果请求类型是 application/x-www-form-urlencoded时解析form数据
	// req.ParseForm()
	// fmt.Println(req.PostForm)
	// fmt.Println(req.PostForm.Get("user"), req.PostForm.Get("pass"))

	// 如果数据格式是json
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("readAll failed, err:%", err)
		return
	}
	fmt.Println(string(b))
	resp := `{"status":"ok"}`
	rw.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/products", getProducts)
	fmt.Println("server start and listen on 8090")
	http.ListenAndServe("0.0.0.0:8090", nil)
}
