package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
)

func GetDataUser(respw http.ResponseWriter, req *http.Request) {
    // Decode token untuk mendapatkan informasi pengguna
    payload, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Token Tidak Valid"
        respn.Info = at.GetSecretFromHeader(req)
        respn.Location = "Decode Token Error"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }

    // Cari data pengguna berdasarkan nomor telepon (phonenumber) yang terdapat di token
    docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Pengguna Tidak Ditemukan"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusNotFound, respn)
        return
    }

    // Kembalikan data pengguna yang ditemukan
    at.WriteJSON(respw, http.StatusOK, docuser)
}

func PostDataUser(respw http.ResponseWriter, req *http.Request) {
    // Decode body input
    var usr model.Userdomyikado
    err := json.NewDecoder(req.Body).Decode(&usr)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Body tidak valid"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    // Validasi input user
    if usr.Name == "" || usr.PhoneNumber == "" || usr.Email == "" {
        var respn model.Response
        respn.Status = "Error: Isian tidak lengkap"
        respn.Response = "Mohon isi lengkap Name, PhoneNumber, dan Email"
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    // Masukkan data user baru ke database
    idusr, err := atdb.InsertOneDoc(config.Mongoconn, "user", usr)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Gagal Insert Database"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusInternalServerError, respn)
        return
    }

    usr.ID = idusr // Assign ID yang dihasilkan dari database
    at.WriteJSON(respw, http.StatusOK, usr)
}

func UpdateDataUser(respw http.ResponseWriter, req *http.Request) {
    // Mengambil token dari header dan log untuk debugging
    token := at.GetLoginFromHeader(req)
    fmt.Println("Token received:", token)

    // Decode token untuk mendapatkan informasi pengguna
    payload, err := watoken.Decode(config.PublicKeyWhatsAuth, token)
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Token Tidak Valid"
        respn.Info = at.GetSecretFromHeader(req)
        respn.Location = "Decode Token Error"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }

    // Mengambil data user yang akan diupdate
    var usr model.Userdomyikado
    err = json.NewDecoder(req.Body).Decode(&usr)
    if err != nil {
        var respn model.Response
        respn.Status = "Error : Body tidak valid"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    // Validasi bahwa nama tidak kosong
    if usr.Name == "" {
        var respn model.Response
        respn.Status = "Error: Nama tidak boleh kosong"
        at.WriteJSON(respw, http.StatusBadRequest, respn)
        return
    }

    // Log untuk memastikan data user yang diterima benar
    fmt.Println("User data received:", usr)

    // Update data nama user di database berdasarkan phone number yang ada di token
    filter := bson.M{"phonenumber": payload.Id}
    
    // Pastikan kita hanya melakukan update pada field tertentu, dalam hal ini nama
    update := bson.M{"$set": bson.M{"name": usr.Name}}

    // Log untuk melihat filter dan update yang digunakan dalam query
    fmt.Println("Filter:", filter)
    fmt.Println("Update:", update)

    // Menggunakan UpdateOne untuk mengupdate nama pengguna
    _, err = atdb.UpdateOneDoc(config.Mongoconn, "user", filter, update)
    if err != nil {
        var respn model.Response
        respn.Status = "Error: Gagal Update Database"
        respn.Response = err.Error()
        at.WriteJSON(respw, http.StatusInternalServerError, respn)
        return
    }

    // Log untuk memastikan update berhasil
    fmt.Println("Database updated successfully.")

    // Kembalikan data pengguna yang telah diperbarui
    at.WriteJSON(respw, http.StatusOK, usr)
}

