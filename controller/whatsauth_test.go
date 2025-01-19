package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/whatsauth/itmodel"
// )

// // Mock untuk WhatsAuthService
// type MockWhatsAuthService struct {
// 	mock.Mock
// }

// func (m *MockWhatsAuthService) GetAppProfile(waphonenumber string, mongoconn interface{}) (*itmodel.AppProfile, error) {
// 	args := m.Called(waphonenumber, mongoconn)
// 	return args.Get(0).(*itmodel.AppProfile), args.Error(1)
// }

// func (m *MockWhatsAuthService) WebHook(profile *itmodel.AppProfile, msg itmodel.IteungMessage, mongoconn interface{}) (itmodel.Response, error) {
// 	args := m.Called(profile, msg, mongoconn)
// 	return args.Get(0).(itmodel.Response), args.Error(1)
// }

// func TestPostInboxNomor(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// Mocks
// 	mockWhatsAuth := NewMockWhatsAuthService(ctrl)
// 	mockWhatsAuth.EXPECT().GetAppProfile(gomock.Any(), gomock.Any()).Return(&itmodel.AppProfile{}, nil)
// 	mockWhatsAuth.EXPECT().WebHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(itmodel.Response{}, nil)

// 	// Setup request
// 	body := `{"msg":"test message"}`
// 	req, err := http.NewRequest("POST", "/inbox/nomor", bytes.NewBufferString(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req.Header.Set("Secret", "some-secret")

// 	// Mocking response writer
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(PostInboxNomor)
// 	handler.ServeHTTP(rr, req)

// 	// Check status code
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// Check response body
// 	var resp itmodel.Response
// 	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.Response != "Success" {
// 		t.Errorf("expected success, got %v", resp.Response)
// 	}
// }
