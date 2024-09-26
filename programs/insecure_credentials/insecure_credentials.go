package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"
)

const sampleCertPEM = `-----BEGIN CERTIFICATE-----
MIIDCjCCAfKgAwIBAgIQIj4BuOtQRWxvUA4CUaL+WjANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTE4MDIyMjEzNDA1NFoYDzIxMzIwMzIzMDU0
MDU0WjASMRAwDgYDVQQKEwdBY21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEA0f+rxvg+P/YxJ9Rnj4qPypexre9OAwHfYIfDtBwPffSNhmaJa6Ir
JmPDAfrkGmAu8E+1EJMRge7R4js+y7lj/nxSTHQW4ixXWYNaHrXB8T2ty+dW2T+t
TWagtBkgdZqC+t3AloRtDJBIFKXcd6yHA9q9vj/KRtnafTPjDYD+m4obR5vhkFYm
5oJJoLkcuZ8hGr3MdzHFMIPOJ5Bm5YBY3z4TLqGnmDqhL3pqNHW0xHP7wGEJOTal
I/3OqRthAkLLMwUCHQcpLt1j2jTbavodUSr4ibNXTn5L1ynRGtozb2iE+4bZlRQZ
oR0Q32XxPQ+vkKtatgXS7E6yiq/vUc88hQIDAQABo1owWDAOBgNVHQ8BAf8EBAMC
AqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zAgBgNVHREE
GTAXgglsb2NhbGhvc3SHBAAAAACHBH8AAAEwDQYJKoZIhvcNAQELBQADggEBAJgo
hrLJDKN9VXh6EXYtaeMxRVEINt+swrXLoxNcNmRXZb5vX11yX9uHWCcIaOHZM4c6
+ZZe6gtdTGswrzl7vB5RJ5ZJEypj0MhAvH/PN0J9W0gXYbxzI839RQ2DqNXDjU7I
bEDlKBSSmFb0TjXTuXhHKyviLETAbf143Zb7M1i9L+U5fiPaq2Zt07NX6d2SYeMd
7udXyv/WhWfXKYj2Hoa8sKfcNr2e68IkbD6i1j9zXSbOMfvs1JZgryGqNIoGDOPz
+M3QhvvuiYJCSoOhDph0pNoVeH4NtaVwqPe7qMPnim11CGQSfjzxmZMFqsoJIsRe
lig/ubNJZbC6oA1X+t4=
-----END CERTIFICATE-----`

func main() {
	// Decode and parse the PEM certificate
	cert, err := parseCert(sampleCertPEM)
	if err != nil {
		log.Fatalf("Couldn't parse the certificate: %v", err)
	}
	printCertInfo(cert)
	validateCert(cert)
}

// Parses a PEM-encoded certificate
func parseCert(pemData string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("couldn't decode PEM block")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}
	return cert, nil
}

// Print basic certificate information
func printCertInfo(cert *x509.Certificate) {
	fmt.Println("Certificate Details:")
	fmt.Printf("  Issuer: %s\n", cert.Issuer)
	fmt.Printf("  Subject: %s\n", cert.Subject)
	fmt.Printf("  Valid From: %s\n", cert.NotBefore)
	fmt.Printf("  Valid Until: %s\n", cert.NotAfter)
	fmt.Printf("  Serial Number: %s\n", cert.SerialNumber)
	fmt.Printf("  Public Key Algorithm: %s\n", cert.PublicKeyAlgorithm)
	fmt.Printf("  Signature Algorithm: %s\n", cert.SignatureAlgorithm)
}

// Check if the certificate is currently valid
func validateCert(cert *x509.Certificate) {
	now := time.Now()
	switch {
	case now.Before(cert.NotBefore):
		fmt.Println("Certificate is not valid yet. Please contact us.")
	case now.After(cert.NotAfter):
		fmt.Println("Certificate has expired. Please update ASAP.")
	default:
		fmt.Println("Certificate is currently valid. No need for update.")
	}
}