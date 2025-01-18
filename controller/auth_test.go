package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gocroot/controller/auth"

// 	"github.com/gocroot/helper/atdb"
// )

// func TestAuth(t *testing.T) {
// 	// Mock request body
// 	request := map[string]string{"token": "dummyToken"}
// 	body, _ := json.Marshal(request)
// 	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewReader(body))
// 	resp := httptest.NewRecorder()

// 	// Mock fungsi
// 	mockVerifyIDToken := func(token, clientID string) (*auth.Payload, error) {
// 		return &auth.Payload{
// 			Claims: map[string]interface{}{
// 				"name":    "Test User",
// 				"email":   "test@example.com",
// 				"picture": "http://example.com/pic.jpg",
// 			},
// 		}, nil
// 	}
// 	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (auth.GoogleCredential, error) {
// 		return auth.GoogleCredential{ClientID: "mocked_client_id"}, nil
// 	}

// 	originalVerifyIDToken := auth.VerifyIDToken
// 	originalGetOneDoc := atdb.GetOneDoc
// 	defer func() {
// 		auth.VerifyIDToken = originalVerifyIDToken
// 		atdb.GetOneDoc = originalGetOneDoc
// 	}()
// 	auth.VerifyIDToken = mockVerifyIDToken
// 	atdb.GetOneDoc = mockGetOneDoc

// 	// Panggil fungsi
// 	Auth(resp, req)

// 	// Validasi hasil
// 	if resp.Code != http.StatusOK {
// 		t.Errorf("expected status OK, got %d", resp.Code)
// 	}
// }