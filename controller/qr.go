package controller

import (
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

	helper.WriteJSON(respw, http.StatusOK, user)
}

// Create User
func PostUser(respw http.ResponseWriter, req *http.Request) {
	var newUser model.Users
	if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	// Check for duplicate email
	count, err := atdb.GetCountDoc(config.Mongoconn, "users", bson.M{"email": newUser.Email})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	if count > 0 {
		helper.WriteJSON(respw, http.StatusConflict, "Email already exists")
		return
	}

	if _, err := atdb.InsertOneDoc(config.Mongoconn, "users", newUser); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, newUser)
}

// Update User
func UpdateUser(respw http.ResponseWriter, req *http.Request) {
	var updateUser model.Users
	if err := json.NewDecoder(req.Body).Decode(&updateUser); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	if updateUser.ID == primitive.NilObjectID {
		helper.WriteJSON(respw, http.StatusBadRequest, "User ID is required")
		return
	}

	filter := bson.M{"_id": updateUser.ID}
	update := bson.M{
		"$set": bson.M{
			"name":       updateUser.Name,
			"profilePic": updateUser.ProfilePic,
			"updatedAt":  time.Now(),
		},
	}

	if _, err := atdb.UpdateOneDoc(config.Mongoconn, "users", filter, update); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, "User updated successfully")
}

// Delete User
func DeleteUser(respw http.ResponseWriter, req *http.Request) {
	var user model.Users
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := atdb.DeleteOneDoc(config.Mongoconn, "users", bson.M{"_id": user.ID}); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, "User deleted successfully")
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

	if err := atdb.DeleteOneDoc(config.Mongoconn, "qrhistory", bson.M{"_id": qr.ID}); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(respw, http.StatusOK, "QR History deleted successfully")
}