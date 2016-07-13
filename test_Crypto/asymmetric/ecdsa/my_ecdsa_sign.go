//https://www.socketloop.com/tutorials/golang-example-for-ecdsa-elliptic-curve-digital-signature-algorithm-functions

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"hash"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"test/test_Crypto/asymmetric/ecdsa/fixed_Ecdsa"
)

func main() {
	spew.Config.Indent = "\t"

	privatekey := fixed_Ecdsa.ReturnFixPrivateKey()
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Printf("Private Key :")
	spew.Dump(privatekey)

	// Sign ecdsa style
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
	signhash := h.Sum(nil)

	r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if serr != nil {
		fmt.Println(serr)
		os.Exit(1)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Println("Signature :")
	r_cal := big.NewInt(0)
	s_cal := big.NewInt(0)
	// this recalculated code hasn't worked yet
	for idx, value := range signature {
		if idx < len(r.Bytes()) {
			// shift_val = 2^(8*idx)
			shift_val := big.NewInt(0)
			shift_val.Exp(big.NewInt(2), big.NewInt(int64(idx*8)), nil)

			// value_big = value << (8*idx)
			//           = value * 2^(8*idx)
			//           = value * shift_val
			value_big := big.NewInt(0)
			value_big.Mul(big.NewInt(int64(value)), shift_val)

			// total = total + new_val
			(*r_cal).Add(r_cal, value_big)
		} else {
			(*s_cal).Add(big.NewInt(0), big.NewInt(int64(value)))
		}
	}
	fmt.Println("r: ")
	spew.Dump(r)
	fmt.Println("r_cal: ")
	spew.Dump(r_cal)
	fmt.Println("s: ")
	spew.Dump(s)
	fmt.Println("s_cal: ")
	spew.Dump(s_cal)
	spew.Dump(signature)

	// Verify
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	fmt.Println("Verify status = ", verifystatus) // should be true

	// Write to file and test
	type path_value struct {
		path   string
		number *big.Int
	}
	test_cases := []path_value{
		{fixed_Ecdsa.FILE_R_PATH_TO_SAVE, r},
		{fixed_Ecdsa.FILE_S_PATH_TO_SAVE, s},
	}
	for _, test := range test_cases {
		ioutil.WriteFile(test.path, test.number.Bytes(), 0644)
		signature_read_from_file, _ := ioutil.ReadFile(test.path)
		if bytes.Compare(signature_read_from_file, test.number.Bytes()) != 0 {
			panic("Not equal 1")
		}
		bigInt_fromBytes := big.NewInt(0)
		bigInt_fromBytes.SetBytes(signature_read_from_file)
		if bigInt_fromBytes.Cmp(test.number) != 0 {
			panic("Not equal 2")
		}
	}
}
