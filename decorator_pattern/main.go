package main

import (
	"decorator_pattern/cipher"
	"decorator_pattern/lzw"
	"fmt"
)

type Component interface {
	Operator(string) // function
}

var sentData string

type SendComponent struct {
	// 실질적으로 데이터가 나가는 부분
}

// 데이터 보내기 데코레이터
func (self *SendComponent) Operator(data string) {
	// Send data
	// 나갔다는거 모사
	sentData = data
}

type ZipComponent struct {
	com Component
}

// 압축 데코레이터
func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	// 에러가 없는 경우
	self.com.Operator(string(zipData))
}

// 암호화
type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(unzipData))
}

var recvData string

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	sender := &EncryptComponent{key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{}}}

	sender.Operator("Hello World")

	fmt.Println(sentData) // 결과: Ǽe2��ɺQ���"�G��4p�I��N�xBp�am\�ޔ���lV��

	// 압축 풀고 decrypt 적용해 결과 확인
	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}

	receiver.Operator(sentData)

	fmt.Println(recvData)

}
