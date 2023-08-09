package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/marcosd4h/MDMatador/internal/mdm"
)

type CertContainer struct {
	CertData    []byte
	Certificate *x509.Certificate
	PrivateKey  *rsa.PrivateKey
}

func ParseX509Keypair(rawCertificate []byte, rawPrivateKey []byte) (*x509.Certificate, *rsa.PrivateKey, error) {

	certificate, err := x509.ParseCertificate(rawCertificate)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(rawPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	return certificate, privateKey, nil
}

func GetNewMDMCertificate() (*CertContainer, error) {

	// No certificate in the database, so a new one is generated
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour) // Certificate valid for one year

	// The maximum value a serial number can have is 2^160. However, you could limit this further if required.
	serialNumber := new(big.Int).Lsh(big.NewInt(1), 128) // 2^12

	// Creating a certificate template
	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{mdm.C2ProviderID},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certData, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, err
	}

	// Parsing the newly created certificate and returning it along with the private key
	certificate, err := x509.ParseCertificate(certData)
	if err != nil {
		return nil, err
	}

	newCert := &CertContainer{
		CertData:    certData,
		Certificate: certificate,
		PrivateKey:  privateKey,
	}

	return newCert, nil
}
