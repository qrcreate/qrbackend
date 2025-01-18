package config

// import (
// 	"log"
// 	"os"

// 	"github.com/gocroot/helper/atdb"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // MongoClientTest adalah koneksi MongoDB yang digunakan dalam pengujian
// var MongoClientTest *mongo.Database

// func init() {
// 	// Muat file .env jika ada
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("Error loading .env file")
// 	}

// 	MongoStringTest := os.Getenv("MONGOSTRING_TEST")
// 	if MongoStringTest == "" {
// 		log.Fatal("MONGOSTRING_TEST not set in .env file")
// 	}

// 	MongoinfoTest := atdb.DBInfo{
// 		DBString: MongoStringTest,
// 		DBName:   "testing", // Nama database yang digunakan untuk testing
// 	}

// 	// Menghubungkan ke database MongoDB untuk testing
// 	client, err := atdb.MongoConnect(MongoinfoTest)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to test database: %v", err)
// 	}

// 	// Menyimpan client yang terkoneksi untuk digunakan dalam pengujian
// 	MongoClientTest = client.Database(MongoinfoTest.DBName)

// 	log.Println("Successfully connected to test database")
// }
