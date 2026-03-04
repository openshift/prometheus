# Test certificate generation only

This directory contains tools to regenerate TLS certificates used by **scrape** and **tracing** tests only. It is not built or tested by the Go toolchain (Go ignores `testdata`).

When test certs expire, from the **repository root** run:

```bash
bash scrape/testdata/gencert/generate_certs.sh
```

This updates `scrape/testdata/*.cer`, `*.key` and `tracing/testdata/ca.cer` with 50-year validity certs.
