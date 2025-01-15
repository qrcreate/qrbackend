package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"github.com/kimseokgis/backend-ai/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get All Users
func GetUsers(respw http.ResponseWriter, req *http.Request) {
	users, err := atdb.GetAllDoc[[]model.Users](config.Mongoconn, "users", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	// Hapus field sensitif dari response
	for i := range users {
		users[i].Password = ""
		users[i].PasswordHash = ""
	}

	helper.WriteJSON(respw, http.StatusOK, users)
}


// Get User By ID
func GetOneUser(respw http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "Missing user ID")
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID")
		return
	}

	filter := bson.M{"_id": objID}
	user, err := atdb.GetOneDoc[model.Users](config.Mongoconn, "users", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusNotFound, "User not found")
		return
	}

	// Hapus field sensitif dari response
	user.Password = ""
	user.PasswordHash = ""

	helper.WriteJSON(respw, http.StatusOK, user)
}

// Create User
func PostUser(respw http.ResponseWriter, req *http.Request) {
    var newUser model.Users
    if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, "Error parsing request body: "+err.Error())
        return
    }

    // Validasi input
    if newUser.Email == "" || newUser.Password == "" || newUser.Username == "" {
        helper.WriteJSON(respw, http.StatusBadRequest, "All fields (username, email, password) are required")
        return
    }

    // Hash password
    hashedPassword, err := atdb.HashPass(newUser.Password)
    if err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, "Failed to hash password: "+err.Error())
        return
    }

    // Pastikan password plaintext tidak disimpan
    newUser.Password = "" // Hapus password plaintext
    newUser.PasswordHash = hashedPassword

    // Inisialisasi atribut lain
    newUser.ID = primitive.NewObjectID()
    newUser.CreatedAt = time.Now()
    newUser.UpdatedAt = time.Now()

    // Masukkan ke database
    insertedID, err := atdb.InsertOneDoc(config.Mongoconn, "users", newUser)
    if err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, "Error inserting user data: "+err.Error())
        return
    }

    // Buat response JSON
    response := map[string]interface{}{
        "message":  "User registered successfully",
        "user_id":  insertedID,
        "username": newUser.Username,
        "email":    newUser.Email,
    }
    helper.WriteJSON(respw, http.StatusOK, response)
}

// Update User
func UpdateUser(respw http.ResponseWriter, req *http.Request) {
	// Ambil ID dari query parameter
	id := req.URL.Query().Get("id")
	if id == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "User ID is required")
		return
	}

	// Konversi ID ke ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Ambil data dari body request
	var requestBody map[string]string
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Error decoding request body: "+err.Error())
		return
	}

	// Siapkan perubahan
	updateFields := bson.M{}

	// Perbarui username jika disediakan
	if username, exists := requestBody["username"]; exists && username != "" {
		updateFields["username"] = username
	}

	// Perbarui password jika disediakan
	if password, exists := requestBody["password"]; exists && password != "" {
		hashedPassword, err := atdb.HashPass(password)
		if err != nil {
			helper.WriteJSON(respw, http.StatusInternalServerError, "Failed to hash password: "+err.Error())
			return
		}
		updateFields["passwordhash"] = hashedPassword
	}

	// Tambahkan timestamp pembaruan jika ada field yang diubah
	if len(updateFields) > 0 {
		updateFields["updatedAt"] = time.Now()
	}

	// Jika tidak ada field yang diupdate, kirim error
	if len(updateFields) == 0 {
		helper.WriteJSON(respw, http.StatusBadRequest, "No fields to update")
		return
	}

	// Define filter dan update
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": updateFields,
	}

	// Update dokumen di MongoDB
	result, err := config.Mongoconn.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, "Error updating user: "+err.Error())
		return
	}

	// Periksa apakah dokumen ditemukan
	if result.MatchedCount == 0 {
		helper.WriteJSON(respw, http.StatusNotFound, "User not found")
		return
	}

	// Berikan response dengan informasi perubahan
	response := map[string]interface{}{
		"message": "User updated successfully",
		"changes": updateFields, // Tampilkan field yang diubah
	}
	helper.WriteJSON(respw, http.StatusOK, response)
}

// Delete User
func DeleteUser(respw http.ResponseWriter, req *http.Request) {
	// Ambil ID dari body request
	var requestBody struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Error decoding request body: "+err.Error())
		return
	}

	// Validasi apakah ID diberikan
	if requestBody.ID == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "User ID is required")
		return
	}

	// Konversi ID ke ObjectID
	objID, err := primitive.ObjectIDFromHex(requestBody.ID)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Hapus dokumen berdasarkan ID
	filter := bson.M{"_id": objID}
	result, err := config.Mongoconn.Collection("users").DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, "Error deleting user: "+err.Error())
		return
	}

	// Periksa apakah dokumen ditemukan
	if result.DeletedCount == 0 {
		helper.WriteJSON(respw, http.StatusNotFound, "User not found")
		return
	}

	// Respons sukses
	helper.WriteJSON(respw, http.StatusOK, map[string]string{
		"message": "User deleted successfully",
		"user_id": requestBody.ID,
	})
}

// Get All QR History by User ID
func GetQRHistory(respw http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("userId")
	if userID == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "Missing user ID")
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID")
		return
	}

	filter := bson.M{"userId": objID}
	qrHistory, err := atdb.GetAllDoc[[]model.QrHistory](config.Mongoconn, "qrhistory", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, qrHistory)
}

// Create QR History
func PostQRHistory(respw http.ResponseWriter, req *http.Request) {
	var newQR model.QrHistory
	if err := json.NewDecoder(req.Body).Decode(&newQR); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	newQR.ID = primitive.NewObjectID()
	newQR.CreatedAt = time.Now()

	if _, err := atdb.InsertOneDoc(config.Mongoconn, "qrhistory", newQR); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, newQR)
}

// Delete QR History
func DeleteQRHistory(respw http.ResponseWriter, req *http.Request) {
	var qr model.QrHistory
	if err := json.NewDecoder(req.Body).Decode(&qr); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	// Ambil kedua nilai yang dikembalikan oleh DeleteOneDoc
	_, err := atdb.DeleteOneDoc(config.Mongoconn, "qrhistory", bson.M{"_id": qr.ID})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, "QR History deleted successfully")
}
