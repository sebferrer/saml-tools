package saml

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/beevik/etree"
	"github.com/crewjam/saml/xmlenc"
)

func MustParsePrivateKey(pemStr []byte) *rsa.PrivateKey {
	b, _ := pem.Decode(pemStr)
	if b == nil {
		panic("cannot parse PEM")
	}
	k, err := x509.ParsePKCS1PrivateKey(b.Bytes)
	if err != nil {
		panic(err)
	}
	return k
}

func GetAndDecryptAssertion(samlResponse string, key *rsa.PrivateKey) (string, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromBytes([]byte(samlResponse))
	if err != nil {
		return "", err
	}

	responseEl := doc.Root()
	el := responseEl.FindElement("//EncryptedAssertion/EncryptedData")
	plaintextAssertion, err := xmlenc.Decrypt(key, el)

	return string(plaintextAssertion), err
}
