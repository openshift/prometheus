# Scrape testdata

TLS certificates in this directory are used by `scrape/target_test.go`.

## Regenerating the CA certificate (when it expires)

The CA cert (`ca.cer`) may expire and cause `TestNewHTTPCACert` (and other TLS tests) to fail. Regenerate it with the **same key** so existing server/client certs still verify:

```bash
# From repository root, using absolute paths:
openssl req -new -x509 -key scrape/testdata/ca.key -out scrape/testdata/ca.cer \
  -days 3650 \
  -subj "/C=XX/L=Default City/O=Default Company Ltd/CN=Prometheus Test CA"
```

Alternatively, from this directory: `go run gen_ca.go` (requires being in `scrape/testdata` or repo root).
