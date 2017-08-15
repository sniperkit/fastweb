package crypto

import (
	"crypto/aes"
	"crypto/des"

	. "github.com/tsingson/fastweb/encrypt/cipher"
	"github.com/tsingson/fastweb/encrypt/rsa"
)

/*
介绍:创建默认的AES Cipher,使用ECB工作模式、pkcs57填充,算法秘钥长度128 192 256 位 , 使用秘钥作为初始向量
*/
func NewAES(key []byte) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return NewECBMode().Cipher(block, key[:block.BlockSize()]), err
}

/*
介绍:根据指定的工作模式，创建AESCipher,算法秘钥长度128 192 256 位 , 使用秘钥作为初始向量
*/
func NewAESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:创建默认DESCipher,使用ECB工作模式、pkcs57填充,算法秘钥长度64位 , 使用秘钥作为初始向量
*/
func NewDES(key []byte) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return NewECBMode().Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:根据指定的工作模式，创建DESCipher,算法秘钥长度64位,使用秘钥作为初始向量
*/
func NewDESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:创建RSACipher,默认使用pkcs1 padding,pkcs#1v1.5 加密解密，pkcs#1v1.5签名验证.
*/
func NewRSA(key rsa.RsaKey) (rsa.Cipher, error) {
	padding := rsa.NewPKCS1Padding(key.Modulus())
	cipherMode := rsa.NewPKCS1v15Cipher()
	signMode := rsa.NewPKCS1v15Sign()
	return rsa.NewCipher(key, padding, cipherMode, signMode), nil
}

/*
介绍:根据指定的key,和padding来创建RSACipher
*/
func NewRSAWith(key rsa.RsaKey, padding rsa.Padding, cipherMode rsa.CipherMode, signMode rsa.SignMode) (rsa.Cipher, error) {
	return rsa.NewCipher(key, padding, cipherMode, signMode), nil
}
