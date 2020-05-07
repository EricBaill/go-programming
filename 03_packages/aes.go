// 介绍aes的使用
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type AESManager struct {
	block cipher.Block
}

// 新建aes句柄
func NewAESManager(k []byte) (m *AESManager, err error) {
	m = &AESManager{}
	m.block, err = aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	return m, nil
}

const aesKey = "5797373e8bcbf932"
var plianData string = "hello world"
func main() {
	aesManager, err := NewAESManager([]byte(aesKey))
	if err != nil {
		fmt.Println("NewAESManager error")
		return
	}

	// encode
	base64.enc
	//decode
}
