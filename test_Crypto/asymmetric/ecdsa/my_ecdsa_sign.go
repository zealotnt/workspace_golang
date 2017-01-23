//https://www.socketloop.com/tutorials/golang-example-for-ecdsa-elliptic-curve-digital-signature-algorithm-functions

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/davecgh/go-spew/spew"
	"test/test_Crypto/asymmetric/ecdsa/fixed_Ecdsa"
)

func printHex(mesg string, variable_type string, variable_name string, bytes []byte) {
	fmt.Printf(mesg)

	fmt.Printf("%s %s[] = {\r\n\t\t", variable_type, variable_name)
	for idx, val := range bytes {
		if (idx%8 == 0) && (idx != 0) {
			fmt.Printf("\r\n\t\t")
		}
		fmt.Printf("0x%02x, ", val)
	}

	fmt.Printf("\r\n\t};\r\n\r\n")
}

func recalculateBigInt(bytes []byte) (ret big.Int) {
	big_cal := big.NewInt(0)
	max_len := len(bytes) - 1
	for idx, value := range bytes {
		// shift_val = 2^(8*idx)
		shift_val := big.NewInt(0)
		shift_val.Exp(big.NewInt(2), big.NewInt(int64((max_len-idx)*8)), nil)

		// value_big = value << (8*idx)
		//           = value * 2^(8*idx)
		//           = value * shift_val
		value_big := big.NewInt(0)
		value_big.Mul(big.NewInt(int64(value)), shift_val)

		// total = total + new_val
		(*big_cal).Add(big_cal, value_big)
	}

	return *big_cal
}

func main() {
	spew.Config.Indent = "\t"

	privatekey := fixed_Ecdsa.ReturnFixPrivateKey()
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Printf("Private Key :")
	spew.Dump(privatekey)

	printHex("Pub_X:\r\n\t", "u8", "xq3", pubkey.X.Bytes())
	printHex("Pub_Y:\r\n\t", "u8", "yq3", pubkey.Y.Bytes())

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

	signature_r := r.Bytes()
	signature_s := s.Bytes()

	fmt.Println("Signature :")
	r_cal := recalculateBigInt(signature_r)
	s_cal := recalculateBigInt(signature_s)

	fmt.Printf("r:\r\n\t")
	spew.Dump(r)
	fmt.Printf("r_cal:\r\n\t")
	spew.Dump(r_cal)
	fmt.Printf("signature_r:\r\n\t")
	spew.Dump(signature_r)
	printHex("signature_r_hex_c:\r\n\t", "u8", "r3", signature_r)

	fmt.Printf("s:\r\n\t")
	spew.Dump(s)
	fmt.Printf("s_cal:\r\n\t")
	spew.Dump(s_cal)
	fmt.Printf("signature_s:\r\n\t")
	spew.Dump(signature_s)
	printHex("signature_s_hex_c:\r\n\t", "u8", "s3", signature_s)

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
