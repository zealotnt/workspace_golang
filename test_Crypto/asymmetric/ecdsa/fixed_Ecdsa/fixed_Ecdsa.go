package fixed_Ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

var (
	FILE_R_PATH_TO_SAVE = "signature_r.temp"
	FILE_S_PATH_TO_SAVE = "signature_s.temp"
)

func ReturnFixPrivateKey() *ecdsa.PrivateKey {
	// P := "115792089210356248762697446949407573530086143415290314195533631308867097853951"
	// N := "115792089210356248762697446949407573529996955224135760342422259061068512044369"
	// B := "41058363725152142129326129780047268409114441015993725554835256314039467401291"
	// Gx := "48439561293906451759052585252797914202762949526041747995844080717082404635286"
	// Gy := "36134250956749795798585127919587881956611106672985015071877198253568414405109"
	// curveParam := elliptic.CurveParams{
	// 	P:       big.NewInt(0),
	// 	N:       big.NewInt(0),
	// 	B:       big.NewInt(0),
	// 	Gx:      big.NewInt(0),
	// 	Gy:      big.NewInt(0),
	// 	BitSize: int(0),
	// 	Name:    string("P-256"),
	// }
	// curveParam.P.SetString(P, 10)
	// curveParam.N.SetString(N, 10)
	// curveParam.B.SetString(B, 10)
	// curveParam.Gx.SetString(Gx, 10)
	// curveParam.Gy.SetString(Gy, 10)

	X := "39941908888743517433885184640505522150797965682106000957378723812504187955732"
	Y := "2873708436525969972034130005080883211040707307475244669200655569463554438226"
	D := "45545850015276229088360632700760929152173237879439452377555454724075191853223"

	ecdsa_PrivateKey := ecdsa.PrivateKey{
		PublicKey: struct {
			elliptic.Curve
			X, Y *big.Int
		}{
			X: big.NewInt(0),
			Y: big.NewInt(0),
		},
		D: big.NewInt(0),
	}
	ecdsa_PrivateKey.PublicKey.Curve = elliptic.P256()
	ecdsa_PrivateKey.PublicKey.X.SetString(X, 10)
	ecdsa_PrivateKey.PublicKey.Y.SetString(Y, 10)
	ecdsa_PrivateKey.D.SetString(D, 10)

	return &ecdsa_PrivateKey
}
