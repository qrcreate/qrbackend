package controller

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gocroot/helper/atdb" // Ensure this is imported
// 	"github.com/gocroot/model"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // Mocking functions directly instead of overwriting `GetAllDoc`, `GetOneDoc`
// var mockGetAllDoc = func(conn interface{}, collection string, filter interface{}) ([]model.Users, error) {
// 	return []model.Users{
// 		{
// 			ID:        primitive.NewObjectID(),
// 			Name:      "Test User",
// 			Email:     "test@example.com",
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}, nil
// }

// var mockGetOneDoc = func(conn interface{}, collection string, filter interface{}) (model.Users, error) {
// 	return model.Users{
// 		ID:        primitive.NewObjectID(),
// 		Name:      "Test User",
// 		Email:     "test@example.com",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}, nil
// }

// var mockInsertOneDoc = func(conn interface{}, collection string, document interface{}) (interface{}, error) {
// 	return nil, nil
// }

// var mockUpdateOneDoc = func(conn interface{}, collection string, filter interface{}, update interface{}) (interface{}, error) {
// 	return nil, nil
// }

// var mockDeleteOneDoc = func(conn interface{}, collection string, filter interface{}) (interface{}, error) {
// 	return nil, nil
// }

// // Tes GetUsers
// func TestGetUsers(t *testing.T) {
// 	// Use the mock directly for atdb.GetAllDoc
// 	originalGetAllDoc := atdb.GetAllDoc[[]model.Users] // Instantiate with the type
// 	atdb.GetAllDoc = mockGetAllDoc
// 	defer func() { atdb.GetAllDoc = originalGetAllDoc }() // Restore original function after the test

// 	req, err := http.NewRequest("GET", "/qr/user", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetUsers)  // Fungsi yang diimpor dari controller

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	var users []model.Users
// 	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
// 		t.Errorf("failed to decode response body: %v", err)
// 	}

// 	if len(users) != 1 {
// 		t.Errorf("expected 1 user, got %v", len(users))
// 	}
// }

// // Tes GetOneUser
// func TestGetOneUser(t *testing.T) {
// 	// Use the mock directly for atdb.GetOneDoc
// 	originalGetOneDoc := atdb.GetOneDoc[model.Users] // Instantiate with the type
// 	atdb.GetOneDoc = mockGetOneDoc
// 	defer func() { atdb.GetOneDoc = originalGetOneDoc }() // Restore original function after the test

// 	id := primitive.NewObjectID().Hex()
// 	req, err := http.NewRequest("GET", "/qr/user/detail?id="+id, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetOneUser)  // Fungsi yang diimpor dari controller
// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	var user model.Users
// 	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
// 		t.Errorf("failed to decode response body: %v", err)
// 	}

// 	if user.Name != "Test User" {
// 		t.Errorf("expected user name 'Test User', got %v", user.Name)
// 	}
// }
