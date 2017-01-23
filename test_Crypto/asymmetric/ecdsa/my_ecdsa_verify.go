package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"test/test_Crypto/asymmetric/ecdsa/fixed_Ecdsa"
)

func main() {
	spew.Config.Indent = "\t"

	privatekey := fixed_Ecdsa.ReturnFixPrivateKey()
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	// Sign ecdsa style
	var h hash.Hash
	h = sha256.New()
	io.WriteString(h, "abc")
	signhash := h.Sum(nil)

	signature_s_test, _ := ioutil.ReadFile(fixed_Ecdsa.FILE_S_PATH_TO_SAVE)
	signature_r_test, _ := ioutil.ReadFile(fixed_Ecdsa.FILE_R_PATH_TO_SAVE)
	s_test := big.NewInt(0)
	r_test := big.NewInt(0)
	s_test.SetBytes(signature_s_test)
	r_test.SetBytes(signature_r_test)

	// Verify
	verifystatus := ecdsa.Verify(&pubkey, signhash, r_test, s_test)
	fmt.Println("signature verify =>", verifystatus) // should be true
}
