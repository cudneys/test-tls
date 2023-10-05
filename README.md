# test-tls

## Overview
test-tls is a small CLI utility to test TLS certificates.  It supports local and remote certs.

## Downloads
Downloads are available in the [releases](https://github.com/cudneys/test-tls/releases) section.

## Usage
```
Tests TLS Certificates

Usage:
  test-tls [flags]

Flags:
  -h, --help              help for test-tls
  -H, --host string       Remote host to test.
  -l, --loglevel string   Sets the log level (default "INFO")
  -p, --port int          The port to test (default 443)
  -P, --protocol string   The protocol to use for the test (tcp or udp) (default "tcp")

```

### Examples

#### Successful test
```
➜  bin test-tls --host www.google.com
INFO[0000] Testing Remote Host                           addr="www.google.com:443" host=www.google.com port=443 proto=tcp
INFO[0000] Connection Established Successfully           cipher_suite=4865 host=www.google.com negotiated_proto= port=443 proto=tcp remote_addr="74.125.136.103:443" version=1.3
INFO[0000] Completed Test Successfully                   elapsed_time=61.454209ms
➜  bin
```

#### Unsuccessful test
```
➜  bin test-tls --host untrusted-root.badssl.com
INFO[0000] Testing Remote Host                           addr="untrusted-root.badssl.com:443" host=untrusted-root.badssl.com port=443 proto=tcp
ERRO[0000] Connection Unsuccessful!                      error="tls: failed to verify certificate: x509: certificate signed by unknown authority" host=untrusted-root.badssl.com port=443 proto=tcp
ERRO[0000] Faulty Cert Info                              cert_num=0 field=Issuer host=untrusted-root.badssl.com port=443 proto=tcp value="CN=BadSSL Untrusted Root Certificate Authority,O=BadSSL,L=San Francisco,ST=California,C=US"
ERRO[0000] Faulty Cert Info                              cert_num=0 field=Expiry host=untrusted-root.badssl.com port=443 proto=tcp value=2025-July-20
ERRO[0000] Faulty Cert Info                              cert_num=0 field=CommonName host=untrusted-root.badssl.com port=443 proto=tcp value="BadSSL Untrusted Root Certificate Authority"
ERRO[0000] Faulty Cert Info                              cert_num=0 field=DNSName host=untrusted-root.badssl.com name_number=0 port=443 proto=tcp value="*.badssl.com"
ERRO[0000] Faulty Cert Info                              cert_num=0 field=DNSName host=untrusted-root.badssl.com name_number=1 port=443 proto=tcp value=badssl.com
ERRO[0000] Faulty Cert Info                              cert_num=1 field=Issuer host=untrusted-root.badssl.com port=443 proto=tcp value="CN=BadSSL Untrusted Root Certificate Authority,O=BadSSL,L=San Francisco,ST=California,C=US"
ERRO[0000] Faulty Cert Info                              cert_num=1 field=Expiry host=untrusted-root.badssl.com port=443 proto=tcp value=2036-July-02
ERRO[0000] Faulty Cert Info                              cert_num=1 field=CommonName host=untrusted-root.badssl.com port=443 proto=tcp value="BadSSL Untrusted Root Certificate Authority"
ERRO[0000] Completed Test With Errors                    elapsed_time=551.54575ms error="tls: failed to verify certificate: x509: certificate signed by unknown authority"
➜  bin
```
