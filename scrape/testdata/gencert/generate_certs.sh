#!/usr/bin/env bash
# Generate TLS certificates for scrape and tracing tests (50-year validity).
# Run from repository root: bash scrape/testdata/gencert/generate_certs.sh
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/../../.." && pwd)"
cd "$ROOT_DIR"

DAYS=$((50 * 365))
SCRAPE_TD=scrape/testdata
TRACING_TD=tracing/testdata
GENCERT_DIR=scrape/testdata/gencert
mkdir -p "$SCRAPE_TD" "$TRACING_TD"

# CA
openssl genrsa -out "$SCRAPE_TD/ca.key" 4096
openssl req -new -x509 -days "$DAYS" -key "$SCRAPE_TD/ca.key" -out "$SCRAPE_TD/ca.cer" \
  -subj "/C=US/O=Prometheus/OU=Prometheus Certificate Authority/CN=Prometheus TLS CA"

# Server (localhost, 127.0.0.1)
openssl genrsa -out "$SCRAPE_TD/server.key" 4096
openssl req -new -key "$SCRAPE_TD/server.key" -out /tmp/server.csr \
  -subj "/C=US/O=Prometheus/CN=localhost"
openssl x509 -req -days "$DAYS" -in /tmp/server.csr -CA "$SCRAPE_TD/ca.cer" -CAkey "$SCRAPE_TD/ca.key" \
  -CAcreateserial -out "$SCRAPE_TD/server.cer" \
  -extfile "$GENCERT_DIR/ext_server.cnf" -extensions v3_req
rm -f /tmp/server.csr

# Client
openssl genrsa -out "$SCRAPE_TD/client.key" 4096
openssl req -new -key "$SCRAPE_TD/client.key" -out /tmp/client.csr \
  -subj "/C=US/O=Prometheus/CN=localhost"
openssl x509 -req -days "$DAYS" -in /tmp/client.csr -CA "$SCRAPE_TD/ca.cer" -CAkey "$SCRAPE_TD/ca.key" \
  -CAcreateserial -out "$SCRAPE_TD/client.cer"
rm -f /tmp/client.csr

# Servername (prometheus.rocks for TestNewHTTPWithServerName)
openssl genrsa -out "$SCRAPE_TD/servername.key" 4096
openssl req -new -key "$SCRAPE_TD/servername.key" -out /tmp/servername.csr \
  -subj "/C=US/O=Prometheus/CN=prometheus.rocks"
openssl x509 -req -days "$DAYS" -in /tmp/servername.csr -CA "$SCRAPE_TD/ca.cer" -CAkey "$SCRAPE_TD/ca.key" \
  -CAcreateserial -out "$SCRAPE_TD/servername.cer" \
  -extfile "$GENCERT_DIR/ext_servername.cnf" -extensions v3_req
rm -f /tmp/servername.csr

# Copy CA to tracing testdata
cp "$SCRAPE_TD/ca.cer" "$TRACING_TD/ca.cer"

# Clean serial so next run doesn't conflict
rm -f "$SCRAPE_TD/ca.srl"

echo "Generated $SCRAPE_TD/*.cer, *.key and $TRACING_TD/ca.cer"
