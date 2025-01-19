package atdb

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // Pengujian MongoConnect dengan MockDBInfo yang kompatibel
// func TestMongoConnect_Success(t *testing.T) {
// 	// Setup mock DBInfo untuk pengujian
// 	mockDBInfo := MockDBInfo{
// 		DBString: "mongodb://localhost:27017", // URL yang valid untuk MongoDB
// 		DBName:   "testdb",
// 	}

// 	// Memanggil MongoConnect dengan MockDBInfo
// 	db, err := MongoConnect(mockDBInfo)

// 	// Verifikasi hasil
// 	assert.NoError(t, err)  // Pastikan tidak ada error
// 	assert.NotNil(t, db)    // Pastikan db terhubung dengan benar
// }

// func TestMongoConnect_Fail(t *testing.T) {
// 	// Setup mock DBInfo yang salah
// 	mockDBInfo := MockDBInfo{
// 		DBString: "mongodb://invalid-uri", // URL yang salah
// 		DBName:   "testdb",
// 	}

// 	// Memanggil MongoConnect dengan MockDBInfo dan memastikan error terjadi
// 	db, err := MongoConnect(mockDBInfo)

// 	// Verifikasi apakah terjadi error
// 	assert.Error(t, err) // Pastikan error terjadi
// 	assert.Nil(t, db)    // Pastikan db bernilai nil
// }
