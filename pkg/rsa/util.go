package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
)

func ParseKeycloakRSAPublicKey(base64Str string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode public key")
	}
	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse public key")
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, errors.Wrap(err, "unable to type assert public key to rsa.PublicKey")
	}
	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}
