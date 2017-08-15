package crypto_test

import (
	. "crypto"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/tsingson/fastweb/encrypt/crypto"
	"github.com/tsingson/fastweb/encrypt/rsa"
)

func Test_LoadFromPEMFile(t *testing.T) {

	plant := `rsa rsa `
	key, err := rsa.LoadKeyFromPEMFile(
		`/Users/qinshen/git/linksmart/src/github.com/tsingson/fastweb/encrypt/crypto/rsa_public_key.pem`,
		`/Users/qinshen/git/linksmart/src/github.com/tsingson/fastweb/encrypt/crypto/rsa_private_key.pem`,
		rsa.ParsePKCS8Key)
	if err != nil {
		t.Error(err)
		return
	}
	//
	cipher, err := crypto.NewRSA(key)
	if err != nil {
		t.Error(err)
		return
	}

	enT, err := cipher.Encrypt([]byte(plant))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(enT))

	deT, err := cipher.Decrypt(enT)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(deT))

	signBytes, err := cipher.Sign([]byte(plant), SHA1)
	if err != nil {
		t.Error(err)
		return
	}

	sign := base64.StdEncoding.EncodeToString(signBytes)

	fmt.Println(sign)

	errV := cipher.Verify([]byte(plant), signBytes, SHA1)
	if errV != nil {
		t.Error(errV)
		return
	}

	fmt.Println("verify success")
}
