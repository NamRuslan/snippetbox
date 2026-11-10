package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.namruslan.com/internal/assert"
)

func TestCommonHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	commonHeaders(next).ServeHTTP(rr, r)
	rs := rr.Result()

	assertHeader(t, rs, "Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
	assertHeader(t, rs, "Referrer-Policy", "origin-when-cross-origin")
	assertHeader(t, rs, "X-Content-Type-Options", "nosniff")
	assertHeader(t, rs, "X-Frame-Options", "deny")
	assertHeader(t, rs, "X-XSS-Protection", "0")
	assertHeader(t, rs, "Server", "Go")

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}

func assertHeader(t *testing.T, rs *http.Response, header, expected string) {
	assert.Equal(t, rs.Header.Get(header), expected)
}
