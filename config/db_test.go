package config

import (
	"log"
	"os"

	"github.com/gocroot/helper/atdb"
)

// Koneksi database untuk pengujian
var MongoStringTest string = os.Getenv("MONGOSTRING_TEST")

var MongoinfoTest = atdb.DBInfo{
    DBString: MongoStringTest,
    DBName:   "testing",  
}

// MongoconnTest diekspos untuk digusnakan di luar package
var MongoconnTest, ErrorMongoconnTest = atdb.MongoConnect(MongoinfoTest)

func init() {
    if ErrorMongoconnTest != nil {
        log.Fatalf("Failed to connect to test database: %v", ErrorMongoconnTest)
    } else {
        log.Println("Successfully connected to test database")
    }
}
