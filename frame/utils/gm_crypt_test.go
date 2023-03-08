package utils

import "testing"

func TestGenerateKey(t *testing.T) {

	priKey, pubKey, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("priKey : %s", priKey)
	t.Logf("pubKey : %s", pubKey)

	enstr, err := GmEncode("123", pubKey, 0)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("enstr: %s", enstr)
	destr, err := GmDecode(enstr, priKey, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("destr: %s", destr)

}
