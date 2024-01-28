package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasgrvarela/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	respRecorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(respRecorder, req)

	response := respRecorder.Result()

	assert.Equal(t, response.StatusCode, http.StatusOK)

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
