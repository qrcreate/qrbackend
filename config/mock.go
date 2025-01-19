package config

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

// MockMongoClient menggantikan koneksi MongoDB dalam pengujian
type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) GetAllDoc(dest interface{}, collection string, filter interface{}) error {
	args := m.Called(dest, collection, filter)
	return args.Error(0)
}

func (m *MockMongoClient) GetOneDoc(dest interface{}, collection string, filter interface{}) error {
	args := m.Called(dest, collection, filter)
	return args.Error(0)
}

func (m *MockMongoClient) InsertOneDoc(collection string, doc interface{}) (interface{}, error) {
	args := m.Called(collection, doc)
	return args.Get(0), args.Error(1)
}

func (m *MockMongoClient) ReplaceOneDoc(collection string, filter, doc interface{}) (interface{}, error) {
	args := m.Called(collection, filter, doc)
	return args.Get(0), args.Error(1)
}

// Setup mock untuk pengujian
func setupMock() *MockMongoClient {
	mockClient := new(MockMongoClient)
	// Anda bisa menggunakan mockClient sebagai pengganti koneksi MongoDB asli
	return mockClient
}

func TestMongoConnection(t *testing.T) {
	mockClient := setupMock()

	// Setup mock behavior untuk GetAllDoc
	mockClient.On("GetAllDoc", mock.Anything, "users", mock.Anything).Return(nil)

	// Gunakan mockClient untuk pengujian lebih lanjut
	// Misalnya, panggil fungsi yang membutuhkan MongoClientTest atau mockClient
	err := mockClient.GetAllDoc(nil, "users", bson.M{})
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	// Verifikasi bahwa mock tersebut digunakan dengan benar
	mockClient.AssertExpectations(t)
}

func TestMongoConnect(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal("Failed to ping MongoDB:", err)
	}
}

func TestGetAllDoc(t *testing.T) {
	mockClient := setupMock()

	mockClient.On("GetAllDoc", mock.Anything, "users", mock.Anything).Return(nil)

	err := mockClient.GetAllDoc(nil, "users", bson.M{})
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	mockClient.AssertExpectations(t)
}

func TestGetOneDoc(t *testing.T) {
	mockClient := setupMock()

	mockClient.On("GetOneDoc", mock.Anything, "users", mock.Anything).Return(nil)

	err := mockClient.GetOneDoc(nil, "users", bson.M{})
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	mockClient.AssertExpectations(t)
}

func TestInsertOneDoc(t *testing.T) {
	mockClient := setupMock()

	mockClient.On("InsertOneDoc", "users", mock.Anything).Return("mocked_id", nil)

	id, err := mockClient.InsertOneDoc("users", bson.M{"name": "test"})
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	if id != "mocked_id" {
		t.Fatalf("Expected id 'mocked_id', but got: %v", id)
	}

	mockClient.AssertExpectations(t)
}

func TestReplaceOneDoc(t *testing.T) {
	mockClient := setupMock()

	mockClient.On("ReplaceOneDoc", "users", mock.Anything, mock.Anything).Return("mocked_id", nil)

	id, err := mockClient.ReplaceOneDoc("users", bson.M{"name": "test"}, bson.M{"name": "updated_test"})
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	if id != "mocked_id" {
		t.Fatalf("Expected id 'mocked_id', but got: %v", id)
	}

	mockClient.AssertExpectations(t)
}
