package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.kieransweeden.dev/internal/assert"
)

func TestPing(t *testing.T) {
	// Create new app instance with structured logger
	// that discards anything written to it, middlewares depend
	// on this existing against the app instance so we include it
	// to avoid causing unrelated panics.
	app := &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	// Start a new TLS test server, Notice that we defer a call to
	// ts.Close() so that the server is shutdown when the test finishes.
	ts := httptest.NewTLSServer(app.routes()) // using all routes so middleware is included too
	defer ts.Close()

	// The network address that the test server is listening on is contained in
	// the ts.URL field.
	// Use this to send a /ping GET request to the TLS test server
	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
