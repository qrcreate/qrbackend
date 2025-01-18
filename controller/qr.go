package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"github.com/gorilla/mux"
	"github.com/kimseokgis/backend-ai/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get All QR History by User ID
func GetQRHistory(respw http.ResponseWriter, req *http.Request) {
    token := at.GetLoginFromHeader(req) 
  
    // Decode the token to get the user information
    payload, err := watoken.Decode(config.PublicKeyWhatsAuth, token)
    if err != nil {
        helper.WriteJSON(respw, http.StatusUnauthorized, "Invalid Token")
        return
    }
  
    // Extract the user ID from the decoded payload (which is a string)
    userIDStr := payload.Id
  
    // Convert string userID to primitive.ObjectID
    userID, err := primitive.ObjectIDFromHex(userIDStr)
    if err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID format")
        return
    }
  
    // MongoDB filter to get QR codes for this user
    filter := bson.M{"userId": userID}
    qrHistory, err := atdb.GetAllDoc[[]model.QrHistory](config.Mongoconn, "qrhistory", filter)
    if err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
        return
    }
  
    helper.WriteJSON(respw, http.StatusOK, qrHistory)
}

// Create QR History
func PostQRHistory(respw http.ResponseWriter, req *http.Request) {
    // Retrieve token from cookie
    token := at.GetLoginFromHeader(req)
  
    // Decode the token to get user info
    payload, err := watoken.Decode(config.PublicKeyWhatsAuth, token)
    if err != nil {
        helper.WriteJSON(respw, http.StatusUnauthorized, "Invalid Token")
        return
    }
  
    // Get userId from the payload (which is a string)
    userIDStr := payload.Id
  
    // Convert string userID to primitive.ObjectID
    userID, err := primitive.ObjectIDFromHex(userIDStr)
    if err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID format")
        return
    }
  
    // Decode the QR data sent in the request body
    var newQR model.QrHistory
    if err := json.NewDecoder(req.Body).Decode(&newQR); err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
        return
    }
  
    // Set userId for the new QR code
    newQR.UserID = userID
    newQR.CreatedAt = time.Now()
  
    // Insert the new QR code into the database
    newQR.ID = primitive.NewObjectID()
    if _, err := atdb.InsertOneDoc(config.Mongoconn, "qrhistory", newQR); err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
        return
    }
  
    helper.WriteJSON(respw, http.StatusOK, newQR)
}

// Update QR
func PutQRHistory(respw http.ResponseWriter, req *http.Request) {
    token := at.GetLoginFromHeader(req)

    // Decode token to get user information
    payload, err := watoken.Decode(config.PublicKeyWhatsAuth, token)
    if err != nil {
        helper.WriteJSON(respw, http.StatusUnauthorized, "Invalid Token")
        return
    }

    // Get userID from the payload
    userIDStr := payload.Id

    // Convert userID from string to primitive.ObjectID
    userID, err := primitive.ObjectIDFromHex(userIDStr)
    if err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID format")
        return
    }

    // Extract the ID from the URL path using mux
    vars := mux.Vars(req)
    id := vars["id"]  // this is where the id will be extracted from the URL

    // Convert the extracted ID string into primitive.ObjectID
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, "Invalid Object ID")
        return
    }

    // Decode QR code data from the body request
    var updatedQR model.QrHistory
    if err := json.NewDecoder(req.Body).Decode(&updatedQR); err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
        return
    }

    // Ensure the QR code belongs to the logged-in user
    if updatedQR.UserID != userID {
        helper.WriteJSON(respw, http.StatusForbidden, "You are not authorized to edit this QR code")
        return
    }

    // Prepare data to update
    updateData := bson.M{
        "name": updatedQR.Name,    // Update only name
        "url":  updatedQR.URL,     // If URL doesn't change, include it as well
        "createdAt": updatedQR.CreatedAt, // If CreatedAt doesn't change, include it as well
    }

    // Update the QR code in the database
    _, err = atdb.UpdateOneDoc(config.Mongoconn, "qrhistory", bson.M{"_id": objectId}, updateData)
    if err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
        return
    }

    // Send success response
    helper.WriteJSON(respw, http.StatusOK, "QR code updated successfully")
}

  
// Delete QR History
func DeleteQRHistory(respw http.ResponseWriter, req *http.Request) {
    token := at.GetLoginFromHeader(req) 
  
    // Decode the token
    payload, err := watoken.Decode(config.PublicKeyWhatsAuth, token)
    if err != nil {
        helper.WriteJSON(respw, http.StatusUnauthorized, "Invalid Token")
        return
    }
  
    // Get userId from the payload
    userIDStr := payload.Id
  
    // Convert string userID to primitive.ObjectID
    userID, err := primitive.ObjectIDFromHex(userIDStr)
    if err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, "Invalid user ID format")
        return
    }
  
    // Decode the QR code ID from the request body
    var qr model.QrHistory
    if err := json.NewDecoder(req.Body).Decode(&qr); err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
        return
    }
  
    // Ensure the QR code belongs to the logged-in user
    if qr.UserID != userID {
        helper.WriteJSON(respw, http.StatusForbidden, "You are not authorized to delete this QR code")
        return
    }
  
    // Delete the QR code
    _, err = atdb.DeleteOneDoc(config.Mongoconn, "qrhistory", bson.M{"_id": qr.ID})
    if err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
        return
    }
  
    helper.WriteJSON(respw, http.StatusOK, "QR History deleted successfully")
}

//download
// func DownloadQR(respw http.ResponseWriter, req *http.Request) {
//     // Extract the QR code ID from the URL path
//     qrIDStr := at.URLParam(req.URL.Path, "/qrcode/") // Ensure this extracts the ID from URL
//     if qrIDStr == "" {
//         helper.WriteJSON(respw, http.StatusBadRequest, "QR code ID is missing from the URL")
//         return
//     }

//     // Convert the extracted ID string into a primitive.ObjectID
//     qrID, err := primitive.ObjectIDFromHex(qrIDStr)
//     if err != nil {
//         helper.WriteJSON(respw, http.StatusBadRequest, "Invalid QR code ID format")
//         return
//     }

//     // Define the file path based on the ID
//     filePath := filepath.Join("path/to/qr/codes", fmt.Sprintf("%s.png", qrID.Hex()))

//     // Open the QR code file
//     file, err := os.Open(filePath)
//     if err != nil {
//         helper.WriteJSON(respw, http.StatusNotFound, "QR code file not found")
//         return
//     }
//     defer file.Close()

//     // Set the headers to serve the file as an attachment
//     respw.Header().Set("Content-Type", "image/png")
//     respw.Header().Set("Content-Disposition", "attachment; filename=qr-code.png")

//     // Stream the file to the response
//     _, err = io.Copy(respw, file)
//     if err != nil {
//         helper.WriteJSON(respw, http.StatusInternalServerError, "Error streaming the file")
//         return
//     }
// }
