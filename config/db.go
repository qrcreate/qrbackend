package config

import (
	"log"
	"os"

	"github.com/gocroot/helper/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "qrcreate",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)

// Koneksi database untuk pengujian
var MongoStringTest string = os.Getenv("MONGOSTRING")

// Koneksi database test (diekspose dengan huruf kapital)
var MongoInfoTest = atdb.DBInfo{
	DBString: MongoStringTest,
	DBName:   "testing",
}

// MongoconnTest diekspos untuk digunakan di luar package
var MongoconnTest, ErrorMongoconnTest = atdb.MongoConnect(MongoInfoTest)

func init() {
	if ErrorMongoconnTest != nil {
		log.Fatalf("Failed to connect to test database: %v", ErrorMongoconnTest)
	} else {
		log.Println("Successfully connected to test database")
	}
}

