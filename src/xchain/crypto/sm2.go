package crypto

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/asn1"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"io"
	"math/big"
)

const (
	BitSize  = 256
	KeyBytes = (BitSize + 7) / 8
)

type Sm2CipherTextType int32

var (
	sm2H                 = new(big.Int).SetInt64(1)
	sm2SignDefaultUserId = []byte{
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
)

var sm2P256V1 P256V1Curve

type P256V1Curve struct {
	*elliptic.CurveParams
	A *big.Int
}

type PublicKey struct {
	X, Y  *big.Int
	Curve P256V1Curve
}

type PrivateKey struct {
	D     *big.Int
	Curve P256V1Curve
}

type sm2Signature struct {
	R, S *big.Int
}

func init() {
	initSm2P256V1()
}

func initSm2P256V1() {
	sm2P, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	sm2A, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFC", 16)
	sm2B, _ := new(big.Int).SetString("28E9FA9E9D9F5E344D5A9E4BCF6509A7F39789F515AB8F92DDBCBD414D940E93", 16)
	sm2N, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	sm2Gx, _ := new(big.Int).SetString("32C4AE2C1F1981195F9904466A39C9948FE30BBFF2660BE1715A4589334C74C7", 16)
	sm2Gy, _ := new(big.Int).SetString("BC3736A2F4F6779C59BDCEE36B692153D0A9877CC62A474002DF32E52139F0A0", 16)
	sm2P256V1.CurveParams = &elliptic.CurveParams{Name: "SM2-P-256-V1"}
	sm2P256V1.P = sm2P
	sm2P256V1.A = sm2A
	sm2P256V1.B = sm2B
	sm2P256V1.N = sm2N
	sm2P256V1.Gx = sm2Gx
	sm2P256V1.Gy = sm2Gy
	sm2P256V1.BitSize = BitSize
}

func GenerateKey(rand io.Reader) (*PrivateKey, *PublicKey, error) {
	priv, x, y, err := elliptic.GenerateKey(sm2P256V1, rand)
	if err != nil {
		return nil, nil, err
	}
	privateKey := new(PrivateKey)
	privateKey.Curve = sm2P256V1
	privateKey.D = new(big.Int).SetBytes(priv)
	publicKey := new(PublicKey)
	publicKey.Curve = sm2P256V1
	publicKey.X = x
	publicKey.Y = y
	return privateKey, publicKey, nil
}

func RawBytesToPublicKey(bytes []byte) (*PublicKey, error) {
	if len(bytes) != KeyBytes*2+1 {
		msg := fmt.Sprintf("Public key raw bytes length must be %d", KeyBytes*2+1)
		return nil, errors.New(msg)
	}
	publicKey := new(PublicKey)
	publicKey.Curve = sm2P256V1
	publicKey.X = new(big.Int).SetBytes(bytes[1 : KeyBytes+1])
	publicKey.Y = new(big.Int).SetBytes(bytes[KeyBytes+1:])
	return publicKey, nil
}

func RawBytesToPrivateKey(bytes []byte) (*PrivateKey, error) {
	if len(bytes) != KeyBytes && len(bytes) != KeyBytes+1 {
		msg := fmt.Sprintf("Private key raw bytes length must be %d or %d", KeyBytes, KeyBytes+1)
		return nil, errors.New(msg)
	}
	privateKey := new(PrivateKey)
	privateKey.Curve = sm2P256V1
	privateKey.D = new(big.Int).SetBytes(bytes)
	return privateKey, nil
}

func nextK(rnd io.Reader, max *big.Int) (*big.Int, error) {
	intOne := new(big.Int).SetInt64(1)
	var k *big.Int
	var err error
	for {
		k, err = rand.Int(rnd, max)
		if err != nil {
			return nil, err
		}
		if k.Cmp(intOne) >= 0 {
			return k, err
		}
	}
}

func getZ(digest hash.Hash, curve *P256V1Curve, pubX *big.Int, pubY *big.Int, userId []byte) []byte {
	digest.Reset()

	userIdLen := uint16(len(userId) * 8)
	var userIdLenBytes [2]byte
	binary.BigEndian.PutUint16(userIdLenBytes[:], userIdLen)
	digest.Write(userIdLenBytes[:])
	if userId != nil && len(userId) > 0 {
		digest.Write(userId)
	}

	digest.Write(curve.A.Bytes())
	digest.Write(curve.B.Bytes())
	digest.Write(curve.Gx.Bytes())
	digest.Write(curve.Gy.Bytes())
	digest.Write(pubX.Bytes())
	digest.Write(pubY.Bytes())
	return digest.Sum(nil)
}

func calculateE(digest hash.Hash, curve *P256V1Curve, pubX *big.Int, pubY *big.Int, userId []byte, src []byte) *big.Int {
	z := getZ(digest, curve, pubX, pubY, userId)

	digest.Reset()
	digest.Write(z)
	digest.Write(src)
	eHash := digest.Sum(nil)
	return new(big.Int).SetBytes(eHash)
}

func MarshalSign(r, s *big.Int) ([]byte, error) {
	result, err := asn1.Marshal(sm2Signature{r, s})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UnmarshalSign(sign []byte) (r, s *big.Int, err error) {
	sm2Sign := new(sm2Signature)
	_, err = asn1.Unmarshal(sign, sm2Sign)
	if err != nil {
		return nil, nil, err
	}
	return sm2Sign.R, sm2Sign.S, nil
}

func SignToRS(priv *PrivateKey, userId []byte, in []byte) (r, s *big.Int, err error) {
	digest := New()
	pubX, pubY := priv.Curve.ScalarBaseMult(priv.D.Bytes())
	if userId == nil {
		userId = sm2SignDefaultUserId
	}
	e := calculateE(digest, &priv.Curve, pubX, pubY, userId, in)

	intZero := new(big.Int).SetInt64(0)
	intOne := new(big.Int).SetInt64(1)
	for {
		var k *big.Int
		var err error
		for {
			k, err = nextK(rand.Reader, priv.Curve.N)
			if err != nil {
				return nil, nil, err
			}
			px, _ := priv.Curve.ScalarBaseMult(k.Bytes())
			r = Add(e, px)
			r = Mod(r, priv.Curve.N)

			rk := new(big.Int).Set(r)
			rk = rk.Add(rk, k)
			if r.Cmp(intZero) != 0 && rk.Cmp(priv.Curve.N) != 0 {
				break
			}
		}

		dPlus1ModN := Add(priv.D, intOne)
		dPlus1ModN = ModInverse(dPlus1ModN, priv.Curve.N)
		s = Mul(r, priv.D)
		s = Sub(k, s)
		s = Mod(s, priv.Curve.N)
		s = Mul(dPlus1ModN, s)
		s = Mod(s, priv.Curve.N)

		if s.Cmp(intZero) != 0 {
			break
		}
	}

	return r, s, nil
}

func Sign(priv *PrivateKey, userId []byte, in []byte) ([]byte, error) {
	r, s, err := SignToRS(priv, userId, in)
	if err != nil {
		return nil, err
	}

	return MarshalSign(r, s)
}

func VerifyByRS(pub *PublicKey, userId []byte, src []byte, r, s *big.Int) bool {
	intOne := new(big.Int).SetInt64(1)
	if r.Cmp(intOne) == -1 || r.Cmp(pub.Curve.N) >= 0 {
		return false
	}
	if s.Cmp(intOne) == -1 || s.Cmp(pub.Curve.N) >= 0 {
		return false
	}

	digest := New()
	if userId == nil {
		userId = sm2SignDefaultUserId
	}
	e := calculateE(digest, &pub.Curve, pub.X, pub.Y, userId, src)

	intZero := new(big.Int).SetInt64(0)
	t := Add(r, s)
	t = Mod(t, pub.Curve.N)
	if t.Cmp(intZero) == 0 {
		return false
	}

	sgx, sgy := pub.Curve.ScalarBaseMult(s.Bytes())
	tpx, tpy := pub.Curve.ScalarMult(pub.X, pub.Y, t.Bytes())
	x, y := pub.Curve.Add(sgx, sgy, tpx, tpy)
	if IsEcPointInfinity(x, y) {
		return false
	}

	expectedR := Add(e, x)
	expectedR = Mod(expectedR, pub.Curve.N)
	return expectedR.Cmp(r) == 0
}

func Verify(pub *PublicKey, userId []byte, src []byte, sign []byte) bool {
	r, s, err := UnmarshalSign(sign)
	if err != nil {
		return false
	}

	return VerifyByRS(pub, userId, src, r, s)
}
