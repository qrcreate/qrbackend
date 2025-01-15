package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocroot/config"
	"go.mongodb.org/mongo-driver/bson"
)

// Setup database pengujian
func setupTestDatabase() {
	// Tambahkan data dummy ke database pengujian
	config.MongoconnTest.Collection("users").InsertOne(context.TODO(), bson.M{
		"username": "testuser",
		"email":    "test@example.com",
	})
}

// Cleanup database pengujian
func cleanupTestDatabase() {
	// Hapus semua data di database pengujian
	config.MongoconnTest.Collection("users").DeleteMany(context.TODO(), bson.M{})
}

// Test GetUsers
func TestGetUsers(t *testing.T) {
	// Setup database pengujian
	setupTestDatabase()
	defer cleanupTestDatabase()

	// Buat request dan response recorder
	req, err := http.NewRequest("GET", "/qr/user", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(rr, req)

	// Validasi response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Validasi body response
	expected := `[{"username":"testuser","email":"test@example.com"}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
