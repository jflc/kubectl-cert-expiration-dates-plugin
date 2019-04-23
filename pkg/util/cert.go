package util

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParseCertificate(data []byte) (*x509.Certificate, error) {
	var result *x509.Certificate
	var err error

	block, _ := pem.Decode(data)
	if block == nil {
		err = fmt.Errorf("Failed to parse certificate PEM")
		return result, err
	}

	return x509.ParseCertificate(block.Bytes)
}
