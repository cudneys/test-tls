package remote

import "testing"

func TestExpired(t *testing.T) {
	err := Test("expired.badssl.com", 443, "tcp")
	if err == nil {
		t.Fatal("expired.badssl.com is expired, but no expiry error was returned")
	}
}

func TestWrongHost(t *testing.T) {
	err := Test("wrong.host.badssl.com", 443, "tcp")
	if err == nil {
		t.Fatal("wrong.host.badssl.com is the wrong hostname, but no host error was returned")
	}
}

func TestSelfSigned(t *testing.T) {
	err := Test("self-signed.badssl.com", 443, "tcp")
	if err == nil {
		t.Fatal("self-signed.badssl.com is self signed, but no error was returned")
	}
}

func TestUntrustedRoot(t *testing.T) {
	err := Test("untrusted-root.badssl.com", 445, "tcp")
	if err == nil {
		t.Fatal("untrusted-root.badssl.com is an untrusted root, but no error was returned")
	}
}
