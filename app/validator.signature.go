package app

import (
	"errors"
)

func (m *Simopi) ValidateSignature() error {
	if !m.Signature.Enabled {
		return nil
	}

	if len(m.Signature.HeaderKey) == 0 {
		return errors.New("signature header key cannot be empty")
	}

	if len(m.Signature.Method) == 0 {
		return errors.New("signature method cannot be empty")
	}

	switch m.Signature.Method {
	case SIGN_METHOD_MD5, SIGN_METHOD_SHA1, SIGN_METHOD_SHA256, SIGN_METHOD_PKCS1V15, SIGN_METHOD_PSS:
	default:
		return errors.New("invalid signature method; must be either [ MD5 | SHA1 | SHA256 | PKCS1V15 | PSS ]")
	}

	// TODO check signature key validity

	return nil
}
