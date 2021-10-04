package main

import (
	"encoding/base64"
	"fmt"
)

func encode() {
	var kalimat string
	fmt.Print("Masukan kalimat ASCII:")
	fmt.Scanln(&kalimat)

	var encode = base64.StdEncoding.EncodeToString([]byte(kalimat))
	fmt.Println("Base64 :", encode)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encode)
	var decode = string(decodeByte)
	fmt.Println("ASCII :", decode)
}

func decode() {
	var kata string
	fmt.Print("Masukan kata Base64:")
	fmt.Scanln(&kata)

	var decodeByte, _ = base64.StdEncoding.DecodeString(kata)
	var decode = string(decodeByte)
	fmt.Println("Base64 :", decode)
}
