package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesEncryptCFB(origData []byte, key, iv []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, len(origData))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted, origData)
	return encrypted
}

func AesDecryptCFB(encrypted []byte, key, iv []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

func main() {
	//AesEncryptCFB()
	//AesDecryptCFB()
}
