package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}
	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	
	//user data
	case method == "GET" && path == "/data/user":
		controller.GetDataUser(w, r)
	//user pendaftaran
	case method == "POST" && path == "/auth/register/users": //mendapatkan email gmail
		controller.RegisterGmailAuth(w, r)
	case method == "POST" && path == "/data/user":
		controller.PostDataUser(w, r)
	case method == "PUT" && path == "/data/user":
		controller.UpdateDataUser(w, r)

	// Google Auth
	case method == "POST" && path == "/auth/users":
		controller.Auth(w, r)
	case method == "POST" && path == "/auth/login":
		controller.GeneratePasswordHandler(w, r)
	case method == "POST" && path == "/auth/verify":
		controller.VerifyPasswordHandler(w, r)
	case method == "POST" && path == "/auth/resend":
		controller.ResendPasswordHandler(w, r)

	//qr
	case method == "GET" && path == "/get/qr":
		controller.GetQRHistory(w, r)
	case method == "POST" && path == "/post/qr":
		controller.PostQRHistory(w, r)
	case method == "PUT" && path == "/put/qr":
		controller.PutQRHistory(w, r)
	case method == "DELETE" && path == "/delete/qr":
		controller.DeleteQRHistory(w, r)

//Register
case method == "POST" && path == "/pdfm/register":
	controller.RegisterHandler(w, r)

//Login
case method == "POST" && path == "/pdfm/login":
	controller.GetUser(w, r)

// Merge PDF Handler
// case method == "POST" && path == "/pdfm/merge":
// 	controller.MergePDFHandler(w, r)

//CRUD 
case method == "GET" && path == "/pdfm/get/users":
	controller.GetUsers(w, r)
case method == "POST" && path == "/pdfm/create/users":
	controller.CreateUser(w, r)
case method == "GET" && path == "/pdfm/getone/users":
	controller.GetOneUser(w, r)
case method == "PUT" && path == "/pdfm/update/users":
	controller.UpdateUser(w, r)
case method == "DELETE" && path == "/pdfm/delete/users":
	controller.DeleteUser(w, r)
	}
}
