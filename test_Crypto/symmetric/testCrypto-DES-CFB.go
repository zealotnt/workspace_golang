package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strconv"
)

// var test_mode = 2 		//use for DES testing
// var test_mode = 1 		//use for my_test_tdes_key
// var test_mode = 0 		//use for default + autogen iv

var test_iv = []byte{0x02, 0x90, 0x9B, 0x1C, 0xFB, 0x50, 0xEA, 0x0F}

var test_tdes_key = [][]byte{
	// 8-byte length key, with K1=K2=K3
	[]byte{0xDE, 0x37, 0x08, 0x5F, 0x7E, 0xDA, 0x37, 0x11, 0xDE, 0x37, 0x08, 0x5F, 0x7E, 0xDA, 0x37, 0x11, 0xDE, 0x37, 0x08, 0x5F, 0x7E, 0xDA, 0x37, 0x11},
	// 16-byte length key, with K1=K3
	[]byte{0xDE, 0x37, 0x08, 0x5F, 0x7E, 0xDA, 0x37, 0x11, 0xFA, 0x2D, 0x9F, 0xE7, 0xE0, 0xFC, 0x29, 0xBE, 0xDE, 0x37, 0x08, 0x5F, 0x7E, 0xDA, 0x37, 0x11},
	// 24-byte length key
	[]byte{0x5D, 0x41, 0x40, 0x2A, 0xBC, 0x4B, 0x2A, 0x76, 0xB9, 0x71, 0x9D, 0x91, 0x10, 0x17, 0xC5, 0x92, 0x28, 0xB4, 0x6E, 0xD3, 0xC1, 0x11, 0xE8, 0x51},
}

var test_des_key = []byte{
	0xDE, 0x37, 0x08, 0x5F, 0x7E, 0xDA, 0x37, 0x11,
}

var string_Mode = []string{
	"Mode default + autogen iv",
	"Mode TDES, my specified key",
	"Mode DES, my specified key",
}

var string_Key = []string{
	"8-byte length key, with K1=K2=K3",
	"16-byte length key, with K1=K3",
	"24-byte length key",
}

func main() {
	test_mode, _ := strconv.Atoi(os.Args[1])
	arg_key, _ := strconv.Atoi(os.Args[2])

	fmt.Println(string_Mode[test_mode])
	if test_mode != 2 && test_mode != 0 {
		fmt.Println(string_Key[arg_key])
	} else if test_mode == 2 {
		fmt.Println("DES key")
	} else {
		fmt.Println("Random key")
	}
	fmt.Println("\r\n\r\n")

	// NewTripleDESCipher can also be used when EDE2 is required by
	// duplicating the first 8 bytes of the 16-byte key.
	ede2Key := []byte("example key 1234")
	plaintext := []byte("Hello world !!!!This is TDES test string")

	fmt.Printf("InputText byte:\t% x\n", plaintext)
	fmt.Printf("InputText str:\t%s\n", plaintext)

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%des.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	var tripleDESKey []byte
	if test_mode == 0 {
		tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
		tripleDESKey = append(tripleDESKey, ede2Key[:8]...)
	} else {
		tripleDESKey = test_tdes_key[arg_key]
	}
	fmt.Printf("TDES Key:\t% x\n", tripleDESKey)

	block, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, len(plaintext))

	var iv []byte
	if test_mode == 0 {
		iv = make([]byte, des.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			panic(err)
		}
	} else {
		iv = test_iv
	}

	stream_enc := cipher.NewCFBEncrypter(block, iv)
	fmt.Printf("InitVector:\t% x\n", iv)
	stream_enc.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("Ciphertext:\t% x\n", ciphertext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	decryptedtext := make([]byte, len(ciphertext))
	stream_dec := cipher.NewCFBDecrypter(block, iv)
	stream_dec.XORKeyStream(decryptedtext, ciphertext)
	fmt.Printf("Decryptedtext:\t%s\n", decryptedtext)
}
