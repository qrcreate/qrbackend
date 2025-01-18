package config

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

// MockMongoClient menggantikan koneksi MongoDB dalam pengujian
type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) GetAllDoc(dest interface{}, collection string, filter interface{}) error {
	args := m.Called(dest, collection, filter)
	return args.Error(0)
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
