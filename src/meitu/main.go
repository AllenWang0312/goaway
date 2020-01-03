package main

import (
	"./encrypt"
	"fmt"
)
func main() {
	//x:=[]byte("世界上最邪恶最专制的现代奴隶制国家--朝鲜")
	//key:=[]byte("hgfedcba87654321")
	//x1:=encryptAES(x,key)
	//fmt.Print(string(x1))
	//x2:=decryptAES(x1,key)
	//fmt.Print(string(x2))

	orig := "3930a176880a6619f0ddb76cab70cda8"

	fmt.Println("原文：", orig)
	encryptCode :=encrypt. AesCBCEncrypt(orig, encrypt.Key)
	fmt.Println("密文：", encryptCode)
	decryptCode := encrypt.AesCBCDecrypt(encryptCode, encrypt.Key)
	fmt.Println("解密结果：", decryptCode)
}
