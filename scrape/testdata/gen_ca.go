// gen_ca.go regenerates ca.cer with extended validity (same key) so TLS tests don't fail when the CA expires.
// Run from scrape/testdata: go run gen_ca.go

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	dir := "scrape/testdata"
	if _, err := os.Stat("ca.key"); err == nil {
		dir = "."
	} else if _, err := os.Stat("scrape/testdata/ca.key"); err == nil {
		dir = "scrape/testdata"
	} else {
		fmt.Fprintf(os.Stderr, "ca.key not found in . or scrape/testdata\n")
		os.Exit(1)
	}
	keyPEM, err := os.ReadFile(dir + "/ca.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read ca.key: %v\n", err)
		os.Exit(1)
	}
	block, _ := pem.Decode(keyPEM)
	if block == nil {
		fmt.Fprintf(os.Stderr, "decode ca.key: no PEM block\n")
		os.Exit(1)
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		key2, e2 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if e2 != nil {
			fmt.Fprintf(os.Stderr, "parse ca.key (PKCS1: %v; PKCS8: %v)\n", err, e2)
			os.Exit(1)
		}
		var ok bool
		key, ok = key2.(*rsa.PrivateKey)
		if !ok {
			fmt.Fprintf(os.Stderr, "ca.key is not RSA\n")
			os.Exit(1)
		}
	}

	notBefore := time.Now().Add(-time.Hour)
	notAfter := time.Now().AddDate(15, 0, 0) // 15 years

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:      []string{"XX"},
			Locality:     []string{"Default City"},
			Organization: []string{"Default Company Ltd"},
			CommonName:   "Prometheus Test CA",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create certificate: %v\n", err)
		os.Exit(1)
	}

	out, err := os.Create(dir + "/ca.cer")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create ca.cer: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()
	if err := pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: certDER}); err != nil {
		fmt.Fprintf(os.Stderr, "write ca.cer: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ca.cer regenerated with validity until", notAfter.Format("2006-01-02"))
}
