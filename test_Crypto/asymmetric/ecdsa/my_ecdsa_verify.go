package main

import (
	"crypto/ecdsa"
	"crypto/md5"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"hash"
	"io"
	"io/ioutil"
	"math/big"
	"test/test_Crypto/asymmetric/ecdsa/fixed_Ecdsa"
)

func main() {
	spew.Config.Indent = "\t"

	privatekey := fixed_Ecdsa.ReturnFixPrivateKey()
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	// Sign ecdsa style
	var h hash.Hash
	h = md5.New()
	io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
	signhash := h.Sum(nil)

	signature_s_test, _ := ioutil.ReadFile(fixed_Ecdsa.FILE_S_PATH_TO_SAVE)
	signature_r_test, _ := ioutil.ReadFile(fixed_Ecdsa.FILE_R_PATH_TO_SAVE)
	s_test := big.NewInt(0)
	r_test := big.NewInt(0)
	s_test.SetBytes(signature_s_test)
	r_test.SetBytes(signature_r_test)

	// Verify
	verifystatus := ecdsa.Verify(&pubkey, signhash, r_test, s_test)
	fmt.Println(verifystatus) // should be true
}
