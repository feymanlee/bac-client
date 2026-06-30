package bac

import "testing"

func TestSignProducesDocumentMD5(t *testing.T) {
	got := Sign("demoApp", "secret", "3", "1710000000123")
	want := "ee14b8e0d62171530f3052d8b778f238"
	if got != want {
		t.Fatalf("Sign() = %q, want %q", got, want)
	}
}

func TestSignedQueryIncludesPublicParameters(t *testing.T) {
	q := signedQuery("ak", "sk", "3", "1710000000123")

	if q.Get("appkey") != "ak" {
		t.Fatalf("appkey = %q", q.Get("appkey"))
	}
	if q.Get("auth_ver") != "3" {
		t.Fatalf("auth_ver = %q", q.Get("auth_ver"))
	}
	if q.Get("nonce") != "1710000000123" {
		t.Fatalf("nonce = %q", q.Get("nonce"))
	}
	if q.Get("s") != Sign("ak", "sk", "3", "1710000000123") {
		t.Fatalf("signature = %q", q.Get("s"))
	}
}
