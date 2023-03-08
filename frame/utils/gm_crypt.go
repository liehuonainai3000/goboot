package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

//国密加解密

// 国密 公钥加密
// mode 加密模式 0:C1C3C2  1:C1C2C3
func GmEncode(endstr string, publickKey string, mode int) (result string, err error) {

	//publickKey := ""
	//publickKey = "-----BEGIN PUBLIC KEY-----\r\n" +
	//	"MFkwEwYHKoZIzj0CAQ3=\r\n" +
	//	"-----END PUBLIC KEY-----"

	var smmode int = sm2.C1C3C2
	if mode == 1 {
		smmode = sm2.C1C2C3
	} else if mode == 0 {
		smmode = sm2.C1C3C2
	} else {
		return "", errors.New("mode error")
	}

	// if !strings.Contains(publickKey, "PUBLIC KEY") {
	// 	publickKey = "-----BEGIN PUBLIC KEY-----\r\n" +
	// 		publickKey + "\r\n" +
	// 		"-----END PUBLIC KEY-----"
	// }

	// d2 := []byte(publickKey)
	// pubMen, err := x509.ReadPublicKeyFromPem(d2)

	pubMen, err := x509.ReadPublicKeyFromHex(publickKey)

	if err != nil {
		println(err.Error())
		return
	}

	msg := []byte(endstr)
	ciphertxt, err := sm2.Encrypt(pubMen, msg, nil, smmode)
	if err != nil {
		println(err.Error())
		return
	}

	result = hex.EncodeToString(ciphertxt)

	return

}

// 国密 私解密
// mode 加密模式 0:C1C3C2  1:C1C2C3
func GmDecode(ciphertxt string, privateKey string, mode int) (result string, err error) {

	//privateKey := ""
	//privateKey = "-----BEGIN PRIVATE KEY-----\r\n" +
	//	"MIGTAgEAMBMGByqGSM49\r\n" +
	//	"-----END PRIVATE KEY-----"

	var smmode = sm2.C1C3C2
	if mode == 1 {
		smmode = sm2.C1C2C3
	} else if mode == 0 {
		smmode = sm2.C1C3C2
	} else {
		return "", errors.New("mode error")
	}
	// if !strings.Contains(privateKey, "PRIVATE KEY") {
	// 	privateKey = "-----BEGIN PRIVATE KEY-----\r\n" +
	// 		privateKey + "\r\n" +
	// 		"-----END PRIVATE KEY-----"
	// }

	// d := []byte(privateKey)

	// privateKeys, err := x509.ReadPrivateKeyFromPem(d, nil)
	privateKeys, err := x509.ReadPrivateKeyFromHex(privateKey)
	if err != nil {
		println(err.Error())
		return
	}

	ciphertxtbyte, err := hex.DecodeString(ciphertxt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ciphertxt3, err := sm2.Decrypt(privateKeys, ciphertxtbyte, smmode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result = string(ciphertxt3)

	return
}

func GenerateKey() (priKey, pubKey string, err error) {
	var priv *sm2.PrivateKey
	priv, err = sm2.GenerateKey(rand.Reader)

	if err != nil {
		return
	}

	priKey = x509.WritePrivateKeyToHex(priv)
	pubKey = x509.WritePublicKeyToHex(&priv.PublicKey)
	// b_pubKey, err := x509.MarshalSm2PublicKey(&priv.PublicKey)

	// if err != nil {
	// 	return
	// }

	// b_priKey, err := x509.MarshalSm2PrivateKey(priv, nil)
	// if err != nil {
	// 	return
	// }

	// priKey = hex.EncodeToString(b_priKey)
	// pubKey = hex.EncodeToString(b_pubKey)

	return
}
