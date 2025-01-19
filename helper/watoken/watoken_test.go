package watoken

import (
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	privkey := os.Getenv("PRIVATEKEY")
	str, _ := EncodeforHours(os.Getenv("PHONENUMBER"), "Qrcreate", privkey, 43830)
	println(str)
	//atr, _ := DecodeGetId("", str)
	//println(atr)

}

func TestDecode(t *testing.T) {
	privkey := os.Getenv("PRIVATEKEY")
	pubkey := os.Getenv("PUBLICKEY")
	encoded, err := EncodeforHours("test_id", "test_alias", privkey, 1)
	if err != nil {
		t.Fatalf("Failed to encode: %v", err)
	}

	payload, err := Decode(pubkey, encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if payload.Id != "test_id" {
		t.Errorf("Expected id 'test_id', got '%s'", payload.Id)
	}

	if payload.Alias != "test_alias" {
		t.Errorf("Expected alias 'test_alias', got '%s'", payload.Alias)
	}
}

func TestEncodeforHours(t *testing.T) {
	privkey := os.Getenv("PRIVATEKEY")
	encoded, err := EncodeforHours("test_id", "test_alias", privkey, 1)
	if err != nil {
		t.Fatalf("Failed to encode: %v", err)
	}

	payload, err := Decode(os.Getenv("PUBLICKEY"), encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if payload.Id != "test_id" {
		t.Errorf("Expected id 'test_id', got '%s'", payload.Id)
	}

	if payload.Alias != "test_alias" {
		t.Errorf("Expected alias 'test_alias', got '%s'", payload.Alias)
	}
}

func TestRandomString(t *testing.T) {
	str := RandomString(10)
	if len(str) != 10 {
		t.Errorf("Expected string of length 10, got %d", len(str))
	}
}
