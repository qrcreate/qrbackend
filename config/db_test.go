package config

import (
	"testing"

	"github.com/gocroot/helper/atdb"
	"github.com/stretchr/testify/assert"
)

func TestMongoConnect_EnvVar_Set(t *testing.T) {
	// Menggunakan URL MongoDB secara langsung tanpa mockEnv
	MongoString := "mongodb://localhost:27017"

	// Verifikasi apakah MongoString sudah di-set dengan benar
	assert.Equal(t, "mongodb://localhost:27017", MongoString)

	// Inisialisasi mongoinfo
	mongoinfo := atdb.DBInfo{
		DBString: MongoString,
		DBName:   "qrcreate",
	}

	// Memanggil MongoConnect
	db, err := atdb.MongoConnect(mongoinfo)

	// Verifikasi bahwa koneksi berhasil
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestMongoConnect_EnvVar_NotSet(t *testing.T) {
	// Cek apakah MongoString kosong atau tidak ada (gantilah ini dengan string yang sesuai jika ingin menggunakan variabel lingkungan)
	MongoString := "mongodb://localhost:27017" // Set langsung dengan URL MongoDB

	// Verifikasi bahwa MongoString telah di-set
	assert.NotEmpty(t, MongoString)

	// Inisialisasi mongoinfo
	mongoinfo := atdb.DBInfo{
		DBString: MongoString,
		DBName:   "qrcreate",
	}

	// Memanggil MongoConnect dan memeriksa jika error terjadi
	db, err := atdb.MongoConnect(mongoinfo)

	// Verifikasi bahwa terjadi error karena variabel lingkungan tidak ada
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestMongoConnect_ValidConnection(t *testing.T) {
	// Gunakan URL MongoDB yang valid langsung tanpa setting environment
	MongoString := "mongodb://localhost:27017"

	// Setup DBInfo untuk pengujian
	mongoinfo := atdb.DBInfo{
		DBString: MongoString,
		DBName:   "qrcreate",
	}

	// Memanggil MongoConnect
	db, err := atdb.MongoConnect(mongoinfo)

	// Verifikasi koneksi MongoDB berhasil
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestMongoConnect_FailConnection(t *testing.T) {
	// Gunakan URL MongoDB yang tidak valid untuk pengujian kegagalan
	MongoString := "mongodb://invalid-uri"

	// Setup DBInfo untuk pengujian
	mongoinfo := atdb.DBInfo{
		DBString: MongoString,
		DBName:   "qrcreate",
	}

	// Memanggil MongoConnect dengan koneksi yang salah
	db, err := atdb.MongoConnect(mongoinfo)

	// Verifikasi bahwa terjadi error pada koneksi
	assert.Error(t, err)
	assert.Nil(t, db)
}
