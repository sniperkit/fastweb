package crypto_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/tsingson/fastweb/encrypt/cipher"
	"github.com/tsingson/fastweb/encrypt/crypto"
)

func Test_DES_ECB(t *testing.T) {

	cipher, err := crypto.NewDES([]byte("Z'{ru/^e"))
	if err != nil {
		t.Error(err)
		return
	}

	plant := `des des `

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CBC ok")
}

func Test_DES_CBC(t *testing.T) {

	mode := cipher.NewCBCMode()
	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode) //创建一个des 加密的builder
	if err != nil {
		t.Error(err)
		return
	}

	plant := `des des `

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)

	fmt.Println(string(pp))

	fmt.Println("Test_DES_ECB ok")
}

func Test_DES_CFB(t *testing.T) {

	mode := cipher.NewCFBMode()

	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `des des `
	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CFB ok")
}

func Test_DES_OFB(t *testing.T) {

	mode := cipher.NewOFBMode()

	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `des des `
	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_OFB ok")
}

func Test_DES_CTR(t *testing.T) {

	mode := cipher.NewCTRMode()

	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `des des `
	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CTR ok")
}
