package vm

import (
	"crypto/sha256"
	"errors"

	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/crypto"
)

type ECDsaCrypto struct {
}

func (c *ECDsaCrypto) Hash160(message []byte) []byte {
	temp, _ := common.ToCodeHash(message)
	return temp.ToArray()
}

func (c *ECDsaCrypto) Hash256(message []byte) []byte {
	return []byte{}
}

func (c *ECDsaCrypto) VerifySignature(message []byte, signature []byte, pubkey []byte) (bool, error) {
	pk, err := crypto.DecodePoint(pubkey)
	if err != nil {
		return false, errors.New("[ECDsaCrypto], crypto.DecodePoint failed.")
	}

	digest := sha256.Sum256(message)
	err = crypto.Verify(*pk, digest[:], signature)
	if err != nil {
		return false, errors.New("[ECDsaCrypto], VerifySignature failed.")
	}

	return true, nil
}
