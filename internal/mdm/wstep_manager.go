package mdm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"math/big"
	"strconv"
	"strings"
	"time"

	"go.mozilla.org/pkcs7"
)

type WSTEPManager struct {
	IdentityCertificate *x509.Certificate
	IdentityFingerprint *string
	identityPrivateKey  *rsa.PrivateKey
	serialNumber        *big.Int
}

func NewCertManager(identityCert *x509.Certificate, identityKey *rsa.PrivateKey) (*WSTEPManager, error) {

	// Sanity checks
	if identityCert == nil || identityKey == nil {
		return nil, errors.New("invalid identity certificate or private key")
	}

	// The maximum value a serial number can have is 2^160. However, you could limit this further if required.
	serialNumber := new(big.Int).Lsh(big.NewInt(1), 128) // 2^12

	// Certificate fingerprint
	fingerprint := sha1.Sum(identityCert.Raw)
	fingerprintHexStr := strings.ToUpper(hex.EncodeToString(fingerprint[:]))

	return &WSTEPManager{
		IdentityCertificate: identityCert,
		identityPrivateKey:  identityKey,
		IdentityFingerprint: &fingerprintHexStr,
		serialNumber:        serialNumber}, nil
}

// SignClientCSR returns a signed certificate from the client certificate signing request and the certificate fingerprint
// subject is the common name of the certificate
// clientCSR is the client certificate signing request
func (c *WSTEPManager) SignClientCSR(subject string, clientCSR *x509.CertificateRequest) ([]byte, string, error) {
	if c == nil {
		return nil, "", errors.New("windows mdm identity keypair was not configured")
	}

	if c.IdentityCertificate == nil || c.identityPrivateKey == nil {
		return nil, "", errors.New("invalid identity certificate or private key")
	}

	// populate the client certificate template
	clientCertificate, err := populateClientCert(subject, c.IdentityCertificate, clientCSR)
	if err != nil {
		return nil, "", fmt.Errorf("failed to populate client certificate: %w", err)
	}

	rawSignedCertDER, err := x509.CreateCertificate(rand.Reader, clientCertificate, c.IdentityCertificate, clientCSR.PublicKey, c.identityPrivateKey)
	if err != nil {
		return nil, "", fmt.Errorf("failed to sign client certificate: %v", err.Error())
	}

	// Generate signed cert fingerprint
	fingerprint := sha1.Sum(rawSignedCertDER)
	fingerprintHex := strings.ToUpper(hex.EncodeToString(fingerprint[:]))

	return rawSignedCertDER, fingerprintHex, nil
}

func (c *WSTEPManager) GetIdentityFingerprint() string {
	if c == nil {
		return ""
	}

	return *c.IdentityFingerprint
}

func (c *WSTEPManager) GetIdentityCert() x509.Certificate {
	if c == nil {
		return x509.Certificate{}
	}

	return *c.IdentityCertificate
}

// Populate MDM client identity certificate
func populateClientCert(subject string, issuerCert *x509.Certificate, csr *x509.CertificateRequest) (*x509.Certificate, error) {
	certRenewalPeriodInSecsInt, err := strconv.Atoi(PolicyCertRenewalPeriodInSecs)
	if err != nil {
		return nil, fmt.Errorf("invalid renewal time: %w", err)
	}

	notBeforeDuration := time.Now().Add(time.Duration(certRenewalPeriodInSecsInt) * -time.Second)
	yearDuration := 365 * 24 * time.Hour

	certSubject := pkix.Name{
		OrganizationalUnit: []string{C2ProviderID},
		CommonName:         subject,
	}

	// The maximum value a serial number can have is 2^160. However, you could limit this further if required.
	maxSerialNumber := new(big.Int).Lsh(big.NewInt(1), 128) // 2^12

	tmpl := &x509.Certificate{
		Subject:            certSubject,
		Issuer:             issuerCert.Issuer,
		Version:            csr.Version,
		PublicKey:          csr.PublicKey,
		PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
		Signature:          csr.Signature,
		SignatureAlgorithm: x509.SHA256WithRSA,
		Extensions:         csr.Extensions,
		ExtraExtensions:    csr.ExtraExtensions,
		IPAddresses:        csr.IPAddresses,
		EmailAddresses:     csr.EmailAddresses,
		DNSNames:           csr.DNSNames,
		URIs:               csr.URIs,
		NotBefore:          notBeforeDuration,
		NotAfter:           notBeforeDuration.Add(yearDuration),
		SerialNumber:       maxSerialNumber,
		KeyUsage:           x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}
	return tmpl, nil
}

// GetClientCSR returns the client certificate signing request from the BinarySecurityToken
func GetClientCSR(binSecTokenData string, tokenType BinSecTokenType) (*x509.CertificateRequest, error) {

	// Checking if this is a valid enroll security token (CSR)
	if (tokenType != MDETokenPKCS7) && (tokenType != MDETokenPKCS10) {
		return nil, fmt.Errorf("token type is not valid for MDM enrollment: %d", tokenType)
	}

	// Decoding the Base64 encoded binary security token to obtain the client CSR bytes
	rawCSR, err := base64.StdEncoding.DecodeString(binSecTokenData)
	if err != nil {
		return nil, fmt.Errorf("decoding the binary security token: %v", err)
	}

	//Sanity checks on binary signature token
	//Sanity checks are done on PKCS10 for the moment
	if tokenType == MDETokenPKCS7 {
		// Parse the CSR in PKCS7 Syntax Standard
		pk7CSR, err := pkcs7.Parse(rawCSR)
		if err != nil {
			return nil, fmt.Errorf("parsing the binary security token: %v", err)
		}

		// Verify the signatures of the CSR PKCS7 object
		err = pk7CSR.Verify()
		if err != nil {
			return nil, fmt.Errorf("verifying CSR data: %v", err)
		}

		// Verify signing time
		currentTime := time.Now()
		if currentTime.Before(pk7CSR.GetOnlySigner().NotBefore) || currentTime.After(pk7CSR.GetOnlySigner().NotAfter) {
			return nil, fmt.Errorf("invalid CSR signing time: %v", err)
		}
	}

	// Decode and verify CSR
	certCSR, err := ParseCertificateRequestFromWindowsDevice(rawCSR)
	if err != nil {
		return nil, fmt.Errorf("problem parsing CSR data: %v", err)
	}

	err = certCSR.CheckSignature()
	if err != nil {
		return nil, fmt.Errorf("invalid CSR signature: %v", err)
	}

	if certCSR.PublicKey == nil {
		return nil, fmt.Errorf("invalid CSR public key: %v", err)
	}

	if len(certCSR.Subject.String()) == 0 {
		return nil, fmt.Errorf("invalid CSR subject: %v", err)
	}

	return certCSR, nil
}
