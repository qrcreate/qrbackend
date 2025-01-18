package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gocroot/config"
// 	"github.com/gocroot/model"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // Mock untuk atdb.MongoConnect dan helper.WriteJSON
// type MockMongoClient struct {
// 	mock.Mock
// }

// func (m *MockMongoClient) GetAllDoc(dest interface{}, collection string, filter interface{}) error {
// 	args := m.Called(dest, collection, filter)
// 	return args.Error(0)
// }

// func (m *MockMongoClient) GetOneDoc(dest interface{}, collection string, filter interface{}) error {
// 	args := m.Called(dest, collection, filter)
// 	return args.Error(0)
// }

// func (m *MockMongoClient) GetCountDoc(collection string, filter interface{}) (int64, error) {
// 	args := m.Called(collection, filter)
// 	return args.Get(0).(int64), args.Error(1)
// }

// func (m *MockMongoClient) InsertOneDoc(collection string, document interface{}) error {
// 	args := m.Called(collection, document)
// 	return args.Error(0)
// }

// func (m *MockMongoClient) UpdateOneDoc(collection string, filter, update interface{}) error {
// 	args := m.Called(collection, filter, update)
// 	return args.Error(0)
// }

// func (m *MockMongoClient) DeleteOneDoc(collection string, filter interface{}) error {
// 	args := m.Called(collection, filter)
// 	return args.Error(0)
// }

// // Setup untuk pengujian
// func setup() *MockMongoClient {
// 	mockClient := new(MockMongoClient)
// 	config.Mongoconn = &mongo.Database{} // Update the assignment to use the correct type
// 	return mockClient
// }

// // Test untuk fungsi GetUsers
// func TestGetUsers(t *testing.T) {
// 	mockClient := setup()

// 	// Simulasi data yang dikembalikan oleh GetAllDoc
// 	users := []model.Users{
// 		{ID: primitive.NewObjectID(), Name: "John Doe", Email: "john@example.com"},
// 	}

// 	// Mock panggilan GetAllDoc
// 	mockClient.On("GetAllDoc", mock.Anything, "users", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*[]model.Users)
// 		*arg = users
// 	})

// 	// Membuat request dan response recorder
// 	req, err := http.NewRequest("GET", "/qr/user", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetUsers)

// 	// Memanggil handler
// 	handler.ServeHTTP(rr, req)

// 	// Memverifikasi hasil response
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "John Doe")
// 	mockClient.AssertExpectations(t)
// }

// // Test untuk fungsi GetOneUser
// func TestGetOneUser(t *testing.T) {
// 	mockClient := setup()

// 	userID := primitive.NewObjectID()
// 	user := model.Users{
// 		ID:    userID,
// 		Name:  "Jane Doe",
// 		Email: "jane@example.com",
// 	}

// 	// Mock panggilan GetOneDoc
// 	mockClient.On("GetOneDoc", mock.Anything, "users", bson.M{"_id": userID}).Return(nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*model.Users)
// 		*arg = user
// 	})

// 	// Membuat request dan response recorder
// 	req, err := http.NewRequest("GET", "/qr/user/detail?id="+userID.Hex(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetOneUser)

// 	// Memanggil handler
// 	handler.ServeHTTP(rr, req)

// 	// Memverifikasi hasil response
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "Jane Doe")
// 	mockClient.AssertExpectations(t)
// }

// // Test untuk PostUser (Create User)
// func TestPostUser(t *testing.T) {
// 	mockClient := setup()

// 	newUser := model.Users{
// 		ID:        primitive.NewObjectID(),
// 		Name:      "New User",
// 		Email:     "newuser@example.com",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Mock GetCountDoc untuk memeriksa apakah email sudah ada
// 	mockClient.On("GetCountDoc", "users", bson.M{"email": newUser.Email}).Return(int64(0), nil)

// 	// Mock InsertOneDoc
// 	mockClient.On("InsertOneDoc", "users", newUser).Return(nil)

// 	// Membuat request dan response recorder
// 	userJSON, _ := json.Marshal(newUser)
// 	req, err := http.NewRequest("POST", "/qr/user", bytes.NewBuffer(userJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(PostUser)

// 	// Memanggil handler
// 	handler.ServeHTTP(rr, req)

// 	// Memverifikasi hasil response
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	mockClient.AssertExpectations(t)
// }
// // Test untuk UpdateUser
// func TestUpdateUser(t *testing.T) {
// 	mockClient := setup()

// 	// User yang ingin diperbarui
// 	userID := primitive.NewObjectID()
// 	updateUser := model.Users{
// 		ID:    userID,
// 		Name:  "Updated Name",
// 		Email: "updated@example.com",
// 	}

// 	// Mock untuk memastikan bahwa filter dan update diterima dengan benar
// 	mockClient.On("UpdateOneDoc", "users", bson.M{"_id": userID}, mock.Anything).Return(nil)

// 	// Membuat request dan response recorder
// 	updateUserJSON, _ := json.Marshal(updateUser)
// 	req, err := http.NewRequest("PUT", "/qr/user", bytes.NewBuffer(updateUserJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UpdateUser)

// 	// Memanggil handler
// 	handler.ServeHTTP(rr, req)

// 	// Memverifikasi hasil response
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "User updated successfully")
// 	mockClient.AssertExpectations(t)
// }
// // Test untuk UpdateUser dengan ID tidak valid

// // Test untuk DeleteUser
// func TestDeleteUser(t *testing.T) {
// 	mockClient := setup()

// 	// User yang ingin dihapus
// 	user := model.Users{
// 		ID:    primitive.NewObjectID(),
// 		Name:  "User to be deleted",
// 		Email: "delete@example.com",
// 	}

// 	// Mock untuk memastikan delete berhasil
// 	mockClient.On("DeleteOneDoc", "users", bson.M{"_id": user.ID}).Return(nil)

// 	// Membuat request dan response recorder
// 	deleteUserJSON, _ := json.Marshal(user)
// 	req, err := http.NewRequest("DELETE", "/qr/user", bytes.NewBuffer(deleteUserJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(DeleteUser)

// 	// Memanggil handler
// 	handler.ServeHTTP(rr, req)

// 	// Memverifikasi hasil response
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "User deleted successfully")
// 	mockClient.AssertExpectations(t)
// }



