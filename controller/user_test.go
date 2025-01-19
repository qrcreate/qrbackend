package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	Id    string
	Alias string
}

func TestGetDataUser(t *testing.T) {
	// Mock request
	req := httptest.NewRequest(http.MethodGet, "/getdatauser", nil)
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	// Mock fungsi Decode dan GetOneDoc
	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
		return model.Userdomyikado{PhoneNumber: "123456789", Name: "Test User"}, nil
	}

	// Ganti fungsi asli dengan mock
	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	defer func() { watoken.Decode = originalDecode; atdb.GetOneDoc = originalGetOneDoc }()
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc

	// Panggil fungsi
	GetDataUser(resp, req)

	// Validasi hasil
	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestPostDataUser(t *testing.T) {
	// Mock request body
	user := model.Userdomyikado{
		Name:        "Test User",
		PhoneNumber: "123456789",
		Email:       "test@example.com",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/postdatauser", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	// Mock fungsi Decode, GetOneDoc, dan InsertOneDoc
	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
		return model.Userdomyikado{}, nil
	}
	mockInsertOneDoc := func(dbConn interface{}, collection string, doc interface{}) (primitive.ObjectID, error) {
		return primitive.NewObjectID(), nil
	}

	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	originalInsertOneDoc := atdb.InsertOneDoc
	defer func() {
		watoken.Decode = originalDecode
		atdb.GetOneDoc = originalGetOneDoc
		atdb.InsertOneDoc = originalInsertOneDoc
	}()
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc
	atdb.InsertOneDoc = mockInsertOneDoc

	// Panggil fungsi
	PostDataUser(resp, req)

	// Validasi hasil
	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestGetDataUserFromApi(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/getdatauserfromapi", nil)
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}

	originalDecode := watoken.Decode
	defer func() { watoken.Decode = originalDecode }()
	watoken.Decode = mockDecode

	GetDataUserFromApi(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestPutTokenDataUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/puttokendatauser", nil)
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
		return model.Userdomyikado{PhoneNumber: "123456789", Name: "Test User"}, nil
	}
	mockReplaceOneDoc := func(dbConn interface{}, collection string, filter, doc interface{}) (int64, error) {
		return 1, nil
	}

	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	originalReplaceOneDoc := atdb.ReplaceOneDoc
	defer func() {
		watoken.Decode = originalDecode
		atdb.GetOneDoc = originalGetOneDoc
		atdb.ReplaceOneDoc = originalReplaceOneDoc
	}()
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc
	atdb.ReplaceOneDoc = mockReplaceOneDoc

	PutTokenDataUser(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestUpdateDataUser(t *testing.T) {
	user := model.Userdomyikado{
		Name: "Updated User",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/updatedatauser", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}
	mockUpdateOneDoc := func(dbConn interface{}, collection string, filter, update interface{}) (int64, error) {
		return 1, nil
	}

	originalDecode := watoken.Decode
	originalUpdateOneDoc := atdb.UpdateOneDoc
	defer func() {
		watoken.Decode = originalDecode
		atdb.UpdateOneDoc = originalUpdateOneDoc
	}()
	watoken.Decode = mockDecode
	atdb.UpdateOneDoc = mockUpdateOneDoc

	UpdateDataUser(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}

func TestPostDataBioUser(t *testing.T) {
	user := model.Userdomyikado{
		Bio: "This is a bio",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/postdatabiouser", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer dummyToken")
	resp := httptest.NewRecorder()

	mockDecode := func(publicKey, token string) (Payload, error) {
		return Payload{Id: "123456789", Alias: "Test User"}, nil
	}
	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
		return model.Userdomyikado{PhoneNumber: "123456789", Name: "Test User"}, nil
	}
	mockReplaceOneDoc := func(dbConn interface{}, collection string, filter, doc interface{}) (int64, error) {
		return 1, nil
	}

	originalDecode := watoken.Decode
	originalGetOneDoc := atdb.GetOneDoc
	originalReplaceOneDoc := atdb.ReplaceOneDoc
	defer func() {
		watoken.Decode = originalDecode
		atdb.GetOneDoc = originalGetOneDoc
		atdb.ReplaceOneDoc = originalReplaceOneDoc
	}()
	watoken.Decode = mockDecode
	atdb.GetOneDoc = mockGetOneDoc
	atdb.ReplaceOneDoc = mockReplaceOneDoc

	PostDataBioUser(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", resp.Code)
	}
}
