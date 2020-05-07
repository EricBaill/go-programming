package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// base64ToString()
	//base64ToImg()
	base64ToImg2()
}

func base64ToString() {
	plainData := "hello golang base64 https://github.com/meetbetter/go-programming \\? +~"

	// 标准base64编码
	encodeString := base64.StdEncoding.EncodeToString([]byte(plainData))
	fmt.Println(encodeString)

	// 标准base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，使用URLEncoding，特殊处理url中的?等字符
	uEnc := base64.URLEncoding.EncodeToString([]byte(plainData))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(uDec))
}

func base64Encode() {
	// base64.StdEncoding.Encode()
}

func base64ToImg2() {
	//读原图片
	src, err := ioutil.ReadFile("a.png")
	if err != nil {
		fmt.Println("ReadFile error")
		return
	}

	//base64压缩
	dst := base64.StdEncoding.EncodeToString(src)

	//解压
	plain, err := base64.StdEncoding.DecodeString(dst)
	if err != nil {
		fmt.Println("DecodeString error",err)
		return
	}
	//写入新文件
	//f, _ := os.OpenFile("b.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//defer f.Close()
	//f.Write(plain)
	ioutil.WriteFile("b.png", plain,os.ModePerm)
}
