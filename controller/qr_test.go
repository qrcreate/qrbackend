package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct Payload untuk digunakan pada mock decode
type Payload struct {
	Id    string
	Alias string
}

func TestGetQRHistory(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/getqrhistory", nil)
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	// Mock function untuk watoken.Decode
	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}

	// Mock function untuk atdb.GetOneDoc
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
		return model.Userdomyikado{ID: primitive.NewObjectID(), PhoneNumber: "123456789"}, nil
	}

	// Mock function untuk atdb.GetAllDoc
	mockGetAllDoc := func(dbConn interface{}, collection string, filter interface{}) ([]model.QrHistory, error) {
		return []model.QrHistory{{Name: "QR1"}, {Name: "QR2"}}, nil
	}

	// Menyimpan nilai asli sebelum penggantian fungsi
	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	originalGetAllDoc := atdb.GetAllDoc
	defer func() {
		// Mengembalikan fungsi seperti semula setelah pengujian
		watoken.Decode = originalDecode
		atdb.GetOneDoc = originalGetOneDoc
		atdb.GetAllDoc = originalGetAllDoc
	}()

	// Menetapkan mock function
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc
	atdb.GetAllDoc = mockGetAllDoc

	// Memanggil fungsi yang akan diuji
	GetQRHistory(resp, req)

	// Validasi hasil
	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestPostQRHistory(t *testing.T) {
	prj := model.QrHistory{Name: "NewQR"}
	body, _ := json.Marshal(prj)
	req := httptest.NewRequest(http.MethodPost, "/postqrhistory", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	// Mock function untuk watoken.Decode
	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}

	// Mock function untuk atdb.InsertOneDoc
	mockInsertOneDoc := func(dbConn interface{}, collection string, doc interface{}) (string, error) {
		return "mocked_id", nil
	}

	// Menyimpan nilai asli sebelum penggantian fungsi
	originalDecode := watoken.Decode
	originalInsertOneDoc := atdb.InsertOneDoc
	defer func() {
		// Mengembalikan fungsi seperti semula setelah pengujian
		watoken.Decode = originalDecode
		atdb.InsertOneDoc = originalInsertOneDoc
	}()

	// Menetapkan mock function
	watoken.Decode = mockDecode
	atdb.InsertOneDoc = mockInsertOneDoc

	// Memanggil fungsi yang akan diuji
	PostQRHistory(resp, req)

	// Validasi hasil
	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestPutQRHistory(t *testing.T) {
	// Mock request body
	qr := model.QrHistory{Name: "UpdatedQR"}
	body, _ := json.Marshal(qr)
	req := httptest.NewRequest(http.MethodPut, "/putqrhistory?id=mocked_id", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	// Mock fungsi untuk watoken.Decode
	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}

	// Mock fungsi untuk atdb.GetOneDoc
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.QrHistory, error) {
		return model.QrHistory{ID: primitive.NewObjectID(), Name: "ExistingQR", Secret: "secret", Owner: model.Userdomyikado{ID: primitive.NewObjectID()}}, nil
	}

	// Mock fungsi untuk atdb.ReplaceOneDoc
	mockReplaceOneDoc := func(dbConn interface{}, collection string, filter, doc interface{}) (int64, error) {
		return 1, nil
	}

	// Menyimpan nilai asli sebelum penggantian fungsi
	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	originalReplaceOneDoc := atdb.ReplaceOneDoc
	defer func() {
		// Mengembalikan fungsi seperti semula setelah pengujian
		watoken.Decode = originalDecode
		atdb.GetOneDoc = originalGetOneDoc
		atdb.ReplaceOneDoc = originalReplaceOneDoc
	}()

	// Menetapkan mock function
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc
	atdb.ReplaceOneDoc = mockReplaceOneDoc

	// Memanggil fungsi yang akan diuji
	PutQRHistory(resp, req)

	// Validasi hasil
	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestDeleteQRHistory(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/deleteqrhistory?id=mocked_id", nil)
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	// Mock function untuk watoken.Decode
	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}

	// Mock function untuk atdb.GetOneDoc
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
		return model.Userdomyikado{ID: primitive.NewObjectID(), PhoneNumber: "123456789"}, nil
	}

	// Mock function untuk atdb.DeleteOneDoc
	mockDeleteOneDoc := func(dbConn interface{}, collection string, filter interface{}) (int64, error) {
		return 1, nil
	}

	// Menyimpan nilai asli sebelum penggantian fungsi
	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	originalDeleteOneDoc := atdb.DeleteOneDoc
	defer func() {
		// Mengembalikan fungsi seperti semula setelah pengujian
		watoken.Decode = originalDecode
		atdb.GetOneDoc = originalGetOneDoc
		atdb.DeleteOneDoc = originalDeleteOneDoc
	}()

	// Menetapkan mock function
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc
	atdb.DeleteOneDoc = mockDeleteOneDoc

	// Memanggil fungsi yang akan diuji
	DeleteQRHistory(resp, req)

	// Validasi hasil
	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}
