package app

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"errors"
)

// TODO: signature validation
func (m *Simopi) CheckSignature(b []byte) error {
	if !m.Signature.Enabled {
		return nil
	}

	switch m.Signature.Method {
	case SIGN_METHOD_MD5:
		// TODO: implement md5 signature check
		return errors.New("md5 signing method are not imeplemented")
	case SIGN_METHOD_SHA1:
		// TODO: implement sha1 signature check
		return errors.New("sha1 signing method are not imeplemented")
	case SIGN_METHOD_SHA256:
		// TODO: implement sha256 signature check
		return errors.New("sha256 signing method are not imeplemented")
	case SIGN_METHOD_PKCS1V15:
		m.CheckPKCS1V15(b)
	case SIGN_METHOD_PSS:
		m.CheckPSS(b)
	}

	return errors.New("no signature type assigned")
}

func (m *Simopi) CheckPKCS1V15(body []byte) error {
	var pubKey *rsa.PublicKey
	switch m.Signature.KeyType {
	case PUB_KEY_TYPE_PKCS1:
		pubKey, _ = x509.ParsePKCS1PublicKey([]byte(m.Signature.PublicKey))
	case PUB_KEY_TYPE_PKIX:
		key, _ := x509.ParsePKIXPublicKey([]byte(m.Signature.PublicKey))
		pubKey, _ = key.(*rsa.PublicKey)
	}
	hashed := sha256.Sum256(body)
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], body)
}

func (m *Simopi) CheckPSS(body []byte) error {
	var pubKey *rsa.PublicKey
	switch m.Signature.KeyType {
	case PUB_KEY_TYPE_PKCS1:
		pubKey, _ = x509.ParsePKCS1PublicKey([]byte(m.Signature.PublicKey))
	case PUB_KEY_TYPE_PKIX:
		key, _ := x509.ParsePKIXPublicKey([]byte(m.Signature.PublicKey))
		pubKey, _ = key.(*rsa.PublicKey)
	}
	hashed := sha256.Sum256(body)
	return rsa.VerifyPSS(pubKey, crypto.SHA256, hashed[:], body, nil)
}
