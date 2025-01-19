package controller

import (
	"encoding/json"
	"net/http"
	"time"

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
    filter := bson.M{
        "owner._id": docuser.ID,
    }

    qrHistory, err := atdb.GetAllDoc[[]model.QrHistory](config.Mongoconn, "qrhistory", filter)
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Data QR tidak ditemukan"
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusNotFound, respn)
        return
    }

    
    // Jika tidak ada QR history
if len(qrHistory) == 0 {
    var respn model.Response
    respn.Status = "Info"  // Menggunakan status "Info" untuk menunjukkan bahwa tidak ada data
    respn.Response = "Anda belum membuat QR, silakan buat QR terlebih dahulu."
    helper.WriteJSON(respw, http.StatusOK, respn)  // Status OK dengan pesan info
    return
}


    // Convert CreatedAt to the desired time format and adjust timezone to Jakarta
    loc, _ := time.LoadLocation("Asia/Jakarta") // Set your desired timezone
    for i := range qrHistory {
        // Adjust the time zone for CreatedAt
        qrHistory[i].CreatedAt = qrHistory[i].CreatedAt.In(loc) // Convert to Jakarta timezone
    }

    // Return the QR history with the formatted time
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

    // Get user data based on phone number from the token payload
    docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Data user tidak ditemukan"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusNotImplemented, respn)
        return
    }

    prj.Owner = docuser
    prj.Secret = watoken.RandomString(48)
    prj.Name = normalize.SetIntoID(prj.Name)

    // Set the current timestamp for CreatedAt
    prj.CreatedAt = time.Now()  // Set current timestamp when creating the QR

    // Check if the QR name already exists
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

    var prj model.QrHistory
    err = json.NewDecoder(req.Body).Decode(&prj)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Body tidak valid"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    id := req.URL.Query().Get("id")
    if id == "" {
        var respn model.Response
        respn.Status = "Error: ID tidak ditemukan di query parameter"
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Invalid ID format"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Data user tidak ditemukan"
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusNotImplemented, respn)
        return
    }

    existingprj, err := atdb.GetOneDoc[model.QrHistory](config.Mongoconn, "qrhistory", primitive.M{"_id": objectId, "owner._id": docuser.ID})
    if err != nil {
        var respn model.Response
        respn.Status = "Error: QR tidak ditemukan"
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusNotFound, respn)
        return
    }

    prj.ID = existingprj.ID
    prj.Secret = existingprj.Secret
    prj.Owner = existingprj.Owner

    // Set CreatedAt to current time when updating
    prj.CreatedAt = time.Now() // Update the time to the current moment

    _, err = atdb.ReplaceOneDoc(config.Mongoconn, "qrhistory", primitive.M{"_id": existingprj.ID}, prj)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Gagal memperbarui database"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusInternalServerError, respn)
        return
    }

    at.WriteJSON(respw, http.StatusOK, prj)
}

func DeleteQRHistory(respw http.ResponseWriter, req *http.Request) {
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

    id := req.URL.Query().Get("id")
    if id == "" {
        var respn model.Response
        respn.Status = "Error: ID tidak ditemukan di query parameter"
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Invalid ID format"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Data user tidak ditemukan"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusNotImplemented, respn)
        return
    }

    existingprj, err := atdb.GetOneDoc[model.QrHistory](config.Mongoconn, "qrhistory", primitive.M{"_id": objectId, "owner._id": docuser.ID})
    if err != nil {
        var respn model.Response
        respn.Status = "Error : QR tidak ditemukan"
        respn.Response = "QR dengan ID tersebut tidak ditemukan atau bukan milik Anda"
        at.WriteJSON(respw, http.StatusNotFound, respn)
        return
    }

	_, err = atdb.DeleteOneDoc(config.Mongoconn, "qrhistory", primitive.M{"_id": existingprj.ID})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Gagal menghapus QR"
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusExpectationFailed, respn)
		return
	}
	
	// Kembalikan data QR yang dihapus sebagai respons
	at.WriteJSON(respw, http.StatusOK, existingprj)
	
}
