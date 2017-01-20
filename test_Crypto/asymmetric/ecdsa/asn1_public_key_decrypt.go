package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"fmt"
)

func main() {
	// CKA_EC_POINT value from PKCS#11 interface
	// p256.BitSize = 256
	CKA_EC_POINT := []byte{0x4, 0x41, 0x4, 0xcc, 0x88, 0xea, 0xc3, 0x55, 0x42, 0x65, 0xf4, 0x90, 0x95, 0x26, 0xb4, 0x3d,
		0x73, 0x28, 0xc0, 0x22, 0xc1, 0xa2, 0xe1, 0x70, 0xfa, 0x12, 0x6e, 0xc, 0xa3, 0x1e, 0x9f, 0x69,
		0x2f, 0x45, 0x79, 0x2b, 0xb8, 0x7e, 0x5a, 0x2a, 0x34, 0xcf, 0x7c, 0x89, 0xb1, 0x75, 0x1e, 0x61,
		0xe8, 0xf3, 0xba, 0x50, 0x87, 0xb0, 0xe2, 0x24, 0xc9, 0xcf, 0x4b, 0x11, 0x2c, 0x1b, 0x28, 0x34,
		0x13, 0xa1, 0xaa}

	var ecp []byte
	_, err := asn1.Unmarshal(CKA_EC_POINT, &ecp)
	if err != nil {
		fmt.Printf("Failed to decode ASN.1 encoded CKA_EC_POINT (%s)", err.Error())
		return
	}

	pubKey, err := getPublic(ecp)
	if err != nil {
		fmt.Printf("Failed to decode public key (%s)", err.Error())
		return
	}

	fmt.Printf("Public key: %#v", pubKey)
}

func getPublic(point []byte) (pub crypto.PublicKey, err error) {
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
	return
}
