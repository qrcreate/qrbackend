package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUsers(t *testing.T) {
	// Simpan koneksi database asli
	originalConn := config.Mongoconn
	// Ganti dengan koneksi database pengujian
	config.Mongoconn = config.MongoconnTest
	defer func() { config.Mongoconn = originalConn }() // Kembalikan koneksi asli setelah tes selesai

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetOneUser(t *testing.T) {
	originalConn := config.Mongoconn
	config.Mongoconn = config.MongoconnTest
	defer func() { config.Mongoconn = originalConn }()

	req, err := http.NewRequest("GET", "/user?id=validObjectId", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOneUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostUser(t *testing.T) {
	originalConn := config.Mongoconn
	config.Mongoconn = config.MongoconnTest
	defer func() { config.Mongoconn = originalConn }()

	newUser := model.Users{
		ID:        primitive.NewObjectID(),
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	payload, _ := json.Marshal(newUser)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateUser(t *testing.T) {
	originalConn := config.Mongoconn
	config.Mongoconn = config.MongoconnTest
	defer func() { config.Mongoconn = originalConn }()

	updateUser := model.Users{
		ID:   primitive.NewObjectID(),
		Name: "Updated Name",
	}
	payload, _ := json.Marshal(updateUser)

	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteUser(t *testing.T) {
	originalConn := config.Mongoconn
	config.Mongoconn = config.MongoconnTest
	defer func() { config.Mongoconn = originalConn }()

	deleteUser := model.Users{
		ID: primitive.NewObjectID(),
	}
	payload, _ := json.Marshal(deleteUser)

	req, err := http.NewRequest("DELETE", "/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func init() {
	fmt.Println("MongoconnTest:", config.MongoconnTest) // DEBUG
}
