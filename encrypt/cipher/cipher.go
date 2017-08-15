package cipher

import (
	. "crypto/cipher"
)

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

func NewBlockCipher(padding Padding, encrypt, decrypt BlockMode) Cipher {
	return &blockCipher{
		encrypt: encrypt,
		decrypt: decrypt,
		padding: padding}
}

type blockCipher struct {
	padding Padding
	encrypt BlockMode
	decrypt BlockMode
}

func (blockCipher *blockCipher) Encrypt(plaintext []byte) []byte {
	plaintext = blockCipher.padding.Padding(plaintext, blockCipher.encrypt.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	blockCipher.encrypt.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func (blockCipher *blockCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	blockCipher.decrypt.CryptBlocks(plaintext, ciphertext)
	plaintext = blockCipher.padding.UnPadding(plaintext)
	return plaintext
}

func NewStreamCipher(encrypt Stream, decrypt Stream) Cipher {
	return &streamCipher{
		encrypt: encrypt,
		decrypt: decrypt}
}

type streamCipher struct {
	encrypt Stream
	decrypt Stream
}

func (streamCipher *streamCipher) Encrypt(plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	streamCipher.encrypt.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}
func (streamCipher *streamCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	streamCipher.decrypt.XORKeyStream(plaintext, ciphertext)
	return plaintext
}
