package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gocroot/config"
// 	"github.com/gocroot/model"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // MockDB is a mock implementation of the DBInterface
// type MockDB struct {
// 	mock.Mock
// }

// func (m *MockDB) InsertOneDoc(conn interface{}, collection string, data interface{}) (primitive.ObjectID, error) {
// 	args := m.Called(conn, collection, data)
// 	return args.Get(0).(primitive.ObjectID), args.Error(1)
// }

// func TestPostDataUser(t *testing.T) {
// 	mockDB := new(MockDB)

// 	config.Mongoconn = mockDB

// 	t.Run("Invalid JSON", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte("invalid json")))
// 		rr := httptest.NewRecorder()

// 		// Execute the handler
// 		PostDataUser(rr, req)

// 		// Assert the status code and response body
// 		assert.Equal(t, http.StatusBadRequest, rr.Code)
// 		var resp model.Response
// 		err := json.NewDecoder(rr.Body).Decode(&resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Error: Body tidak valid", resp.Status)
// 	})

// 	// Test case 2: Missing fields (Name, PhoneNumber, Email)
// 	t.Run("Missing Fields", func(t *testing.T) {
// 		usr := model.Userdomyikado{
// 			Name:        "",
// 			PhoneNumber: "",
// 			Email:       "",
// 		}
// 		body, _ := json.Marshal(usr)
// 		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
// 		rr := httptest.NewRecorder()

// 		// Execute the handler
// 		PostDataUser(rr, req)

// 		// Assert the status code and response body
// 		assert.Equal(t, http.StatusBadRequest, rr.Code)
// 		var resp model.Response
// 		err := json.NewDecoder(rr.Body).Decode(&resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Error: Isian tidak lengkap", resp.Status)
// 	})

// 	// Test case 3: Successful insertion into the database
// 	t.Run("Successful Insertion", func(t *testing.T) {
// 		usr := model.Userdomyikado{
// 			Name:        "John Doe",
// 			PhoneNumber: "1234567890",
// 			Email:       "john@example.com",
// 		}
// 		body, _ := json.Marshal(usr)
// 		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
// 		rr := httptest.NewRecorder()

// 		// Mock database insertion
// 		mockDB.On("InsertOneDoc", mock.Anything, "user", usr).Return(primitive.NewObjectID(), nil)

// 		// Execute the handler
// 		PostDataUser(rr, req)

// 		// Assert the status code and response body
// 		assert.Equal(t, http.StatusOK, rr.Code)
// 		var resp model.Userdomyikado
// 		err := json.NewDecoder(rr.Body).Decode(&resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, usr.Name, resp.Name)
// 		assert.Equal(t, usr.PhoneNumber, resp.PhoneNumber)
// 		assert.Equal(t, usr.Email, resp.Email)
// 	})

// 	// Test case 4: Database insertion failure
// 	t.Run("Database Insertion Failure", func(t *testing.T) {
// 		usr := model.Userdomyikado{
// 			Name:        "John Doe",
// 			PhoneNumber: "1234567890",
// 			Email:       "john@example.com",
// 		}
// 		body, _ := json.Marshal(usr)
// 		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
// 		rr := httptest.NewRecorder()

// 		// Mock database insertion failure
// 		mockDB.On("InsertOneDoc", mock.Anything, "user", usr).Return(primitive.ObjectID{}, fmt.Errorf("database error"))

// 		// Execute the handler
// 		PostDataUser(rr, req)

// 		// Assert the status code and response body
// 		assert.Equal(t, http.StatusInternalServerError, rr.Code)
// 		var resp model.Response
// 		err := json.NewDecoder(rr.Body).Decode(&resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Error: Gagal Insert Database", resp.Status)
// 	})
// }
