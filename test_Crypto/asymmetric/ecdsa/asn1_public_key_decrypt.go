// SOURCE: https://play.golang.org/p/2SKF6YrpEl

package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"fmt"
	"math/big"
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

func main() {
	// CKA_EC_POINT value from PKCS#11 interface
	// p256.BitSize = 256
	CKA_EC_POINT := []byte{
		0x4, 0x41, 0x4, 0xcc, 0x88, 0xea, 0xc3, 0x55, 0x42, 0x65, 0xf4, 0x90, 0x95, 0x26, 0xb4, 0x3d,
		0x73, 0x28, 0xc0, 0x22, 0xc1, 0xa2, 0xe1, 0x70, 0xfa, 0x12, 0x6e, 0xc, 0xa3, 0x1e, 0x9f, 0x69,
		0x2f, 0x45, 0x79, 0x2b, 0xb8, 0x7e, 0x5a, 0x2a, 0x34, 0xcf, 0x7c, 0x89, 0xb1, 0x75, 0x1e, 0x61,
		0xe8, 0xf3, 0xba, 0x50, 0x87, 0xb0, 0xe2, 0x24, 0xc9, 0xcf, 0x4b, 0x11, 0x2c, 0x1b, 0x28, 0x34,
		0x13, 0xa1, 0xaa,
	}

	var ecp []byte
	_, err := asn1.Unmarshal(CKA_EC_POINT, &ecp)
	if err != nil {
		fmt.Printf("Failed to decode ASN.1 encoded CKA_EC_POINT (%s)", err.Error())
		return
	}

	pubKey, X, Y, err := getPublic(ecp)
	if err != nil {
		fmt.Printf("Failed to decode public key (%s)", err.Error())
		return
	}

	fmt.Printf("ECDSA point:\r\n\tvalue=%#v\r\n\tlen=%#v\r\n\r\n", ecp, len(ecp))
	fmt.Printf("Public key:\r\n\t%#v\r\n\r\n", pubKey)
	printHex("X []byte:\r\n\t", "u8", "x_pub", X)
	printHex("Y []byte:\r\n\t", "u8", "y_pub", Y)

	X_recal := big.NewInt(0)
	max_len := len(X) - 1
	for idx, value := range X {
		shift_val := big.NewInt(0)
		shift_val.Exp(big.NewInt(2), big.NewInt(int64((max_len-idx)*8)), nil)

		value_big := big.NewInt(0)
		value_big.Mul(big.NewInt(int64(value)), shift_val)

		X_recal.Add(X_recal, value_big)
	}
	fmt.Printf("X recal=\r\n\t%#v\r\n", X_recal)
	X_recal.SetBytes(X)
	fmt.Printf("X SetBytes=\r\n\t%#v\r\n", X_recal)
}

func getPublic(point []byte) (pub crypto.PublicKey, X []byte, Y []byte, err error) {
	var ecdsaPub ecdsa.PublicKey

	ecdsaPub.Curve = elliptic.P256()
	pointLenght := ecdsaPub.Curve.Params().BitSize/8*2 + 1
	if len(point) != pointLenght {
		err = fmt.Errorf("CKA_EC_POINT (%d) does not fit used curve (%d)", len(point), pointLenght)
		return
	}
	ecdsaPub.X, ecdsaPub.Y = elliptic.Unmarshal(ecdsaPub.Curve, point[:pointLenght])
	if ecdsaPub.X == nil {
		err = fmt.Errorf("Failed to decode CKA_EC_POINT")
		return
	}

	if !ecdsaPub.Curve.IsOnCurve(ecdsaPub.X, ecdsaPub.Y) {
		err = fmt.Errorf("Public key is not on Curve")
		return
	}

	pub = &ecdsaPub
	X = ecdsaPub.X.Bytes()
	Y = ecdsaPub.Y.Bytes()
	return
}
