package at

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPresensiThisMonth(t *testing.T) {
	uri := "" //SRVLookup("mongodb+srv://xx:xxx@cxxx.xxx.mongodb.net/")
	print(uri)
}

func TestGetLoginFromHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("login", "test_login")
	login := GetLoginFromHeader(req)
	if login != "test_login" {
		t.Errorf("expected 'test_login', got '%s'", login)
	}
}

func TestGetSecretFromHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("secret", "test_secret")
	secret := GetSecretFromHeader(req)
	if secret != "test_secret" {
		t.Errorf("expected 'test_secret', got '%s'", secret)
	}
}

func TestWriteJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	content := map[string]string{"message": "hello"}
	WriteJSON(rr, http.StatusOK, content)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"hello"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
