package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/normalize"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"github.com/kimseokgis/backend-ai/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetQRHistory(respw http.ResponseWriter, req *http.Request) {
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Token Tidak Valid"
		respn.Info = at.GetLoginFromHeader(req)
		respn.Location = "Decode Token Error"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}
	// Get user data based on phone number from the token payload
	docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data user tidak ditemukan"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotImplemented, respn)
		return
	}

	// Fetch QR history for this user
	// This query filters for QR created by the user (owner)
	filter := bson.M{
		"owner._id": docuser.ID,  // Fetch QR by owner (owner._id matches docuser ID)
	}

	qrHistory, err := atdb.GetAllDoc[[]model.QrHistory](config.Mongoconn, "qrhistory", filter)
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data QR tidak ditemukan"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotFound, respn)
		return
	}

	// If no QR history found, send a response
	if len(qrHistory) == 0 {
		var respn model.Response
		respn.Status = "Error : Data QR tidak ditemukan"
		respn.Response = "Anda belum membuat QR, silakan buat QR terlebih dahulu."
		helper.WriteJSON(respw, http.StatusNotFound, respn)
		return
	}

	// Return the QR history for the owner
	helper.WriteJSON(respw, http.StatusOK, qrHistory)
}

func PostQRHistory(respw http.ResponseWriter, req *http.Request) {
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Token Tidak Valid"
		respn.Info = at.GetLoginFromHeader(req)
		respn.Location = "Decode Token Error"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}
	var prj model.QrHistory
	err = json.NewDecoder(req.Body).Decode(&prj)
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Body tidak valid"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}
	docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data user tidak di temukan"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusNotImplemented, respn)
		return
	}
	prj.Owner = docuser
	prj.Secret = watoken.RandomString(48)
	prj.Name = normalize.SetIntoID(prj.Name)
	existingprj, err := atdb.GetOneDoc[model.QrHistory](config.Mongoconn, "qrhistory", primitive.M{"name": prj.Name})
	if err != nil {
		idprj, err := atdb.InsertOneDoc(config.Mongoconn, "qrhistory", prj)
		if err != nil {
			var respn model.Response
			respn.Status = "Gagal Insert Database"
			respn.Response = err.Error()
			at.WriteJSON(respw, http.StatusNotModified, respn)
			return
		}
		prj.ID = idprj
		at.WriteJSON(respw, http.StatusOK, prj)
	} else {
		var respn model.Response
		respn.Status = "Error : Name QR sudah ada"
		respn.Response = existingprj.Name
		at.WriteJSON(respw, http.StatusConflict, respn)
		return
	}

}

func PutQRHistory(respw http.ResponseWriter, req *http.Request) {
	// Decode token from header
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "Error: Token Tidak Valid"
		respn.Info = at.GetLoginFromHeader(req)
		respn.Location = "Decode Token Error"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}

	// Decode the QR data from the request body
	var prj model.QrHistory
	err = json.NewDecoder(req.Body).Decode(&prj)
	if err != nil {
		var respn model.Response
		respn.Status = "Error: Body tidak valid"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	// Get the ID from the URL query parameters
	id := req.URL.Query().Get("id")
	if id == "" {
		var respn model.Response
		respn.Status = "Error: ID tidak ditemukan di query parameter"
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	// Convert the ID to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		var respn model.Response
		respn.Status = "Error: Invalid ID format"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	// Get user data from the database
	docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		var respn model.Response
		respn.Status = "Error: Data user tidak ditemukan"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotImplemented, respn)
		return
	}

	// Fetch the existing QR based on ID and user ownership
	existingprj, err := atdb.GetOneDoc[model.QrHistory](config.Mongoconn, "qrhistory", primitive.M{"_id": objectId, "owner._id": docuser.ID})
	if err != nil {
		var respn model.Response
		respn.Status = "Error: QR tidak ditemukan"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotFound, respn)
		return
	}

	// Preserve unmodifiable fields
	prj.ID = existingprj.ID
	prj.Secret = existingprj.Secret
	prj.Owner = existingprj.Owner

	// Update the QR document in the database
	_, err = atdb.ReplaceOneDoc(config.Mongoconn, "qrhistory", primitive.M{"_id": existingprj.ID}, prj)
	if err != nil {
		var respn model.Response
		respn.Status = "Error: Gagal memperbarui database"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}

	// Return the updated QR
	at.WriteJSON(respw, http.StatusOK, prj)
}

func DeleteQRHistory(respw http.ResponseWriter, req *http.Request) {
	// Dekode token dari header permintaan
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Token Tidak Valid"
		respn.Info = at.GetLoginFromHeader(req)
		respn.Location = "Decode Token Error"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}

	// Dekode nama proyek dari body permintaan
	var reqBody struct {
		ProjectName string `json:"project_name"`
	}
	err = json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Body tidak valid"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	// Dapatkan data pengguna berdasarkan ID dari payload token
	docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data user tidak di temukan"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusNotImplemented, respn)
		return
	}

	// Cek apakah proyek dengan nama yang diberikan ada dan dimiliki oleh pengguna
	existingprj, err := atdb.GetOneDoc[model.QrHistory](config.Mongoconn, "qrhistory", primitive.M{"name": reqBody.ProjectName, "owner._id": docuser.ID})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data QR tidak di temukan"
		respn.Response = "Qr dengan nama tersebut tidak ditemukan atau bukan milik Anda"
		at.WriteJSON(respw, http.StatusNotFound, respn)
		return
	}

	// Hapus proyek dari koleksi "project" di MongoDB
	_, err = atdb.DeleteOneDoc(config.Mongoconn, "qrhistory", primitive.M{"_id": existingprj.ID})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Gagal menghapus QR"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusExpectationFailed, respn)
		return
	}

	// Berhasil menghapus proyek
	at.WriteJSON(respw, http.StatusOK, map[string]string{"status": "QR berhasil dihapus"})
}
