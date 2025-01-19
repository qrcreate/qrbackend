package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocroot/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Initialize config or mock the database connection here if necessary
	// You can mock atdb methods or use a test database.
}

// Test RegisterHandler
func TestRegisterHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    model.PdfmUsers
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "Valid Registration",
			requestBody: model.PdfmUsers{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusOK,
			expectedMsg:    "Registrasi berhasil",
		},
		{
			name: "Invalid Registration - Missing Email",
			requestBody: model.PdfmUsers{
				Name:     "John Doe",
				Password: "password123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "Name, Email, dan Password wajib diisi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
			w := httptest.NewRecorder()

			RegisterHandler(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var response map[string]string
			err := json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedMsg, response["message"])
		})
	}
}

// Test CreateUser
func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    model.PdfmUsers
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "Valid User Creation",
			requestBody: model.PdfmUsers{
				Name:     "Jane Doe",
				Email:    "jane@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusOK,
			expectedMsg:    "User created successfully",
		},
		{
			name: "Email Already Exists",
			requestBody: model.PdfmUsers{
				Name:     "Duplicate User",
				Email:    "jane@example.com", // Existing email
				Password: "password123",
			},
			expectedStatus: http.StatusConflict,
			expectedMsg:    "Email already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
			w := httptest.NewRecorder()

			CreateUser(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var response map[string]string
			err := json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedMsg, response["message"])
		})
	}
}

// Test UpdateUser
func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "Valid Update",
			requestBody: struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				ID:       "valid_user_id",
				Name:     "Updated Name",
				Email:    "updated@example.com",
				Password: "newpassword123",
			},
			expectedStatus: http.StatusOK,
			expectedMsg:    "User updated successfully",
		},
		{
			name: "Invalid Update - No Fields",
			requestBody: struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				ID: "valid_user_id",
			},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "No fields to update",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPut, "/update", bytes.NewReader(body))
			w := httptest.NewRecorder()

			UpdateUser(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var response map[string]string
			err := json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedMsg, response["message"])
		})
	}
}

// Test DeleteUser
func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    struct{ ID string `json:"id"` }
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "Valid Delete",
			requestBody: struct{ ID string `json:"id"` }{
				ID: "valid_user_id",
			},
			expectedStatus: http.StatusOK,
			expectedMsg:    "User deleted successfully",
		},
		{
			name: "Invalid ID Format",
			requestBody: struct{ ID string `json:"id"` }{
				ID: "invalid_user_id_format",
			},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "Invalid user ID format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodDelete, "/delete", bytes.NewReader(body))
			w := httptest.NewRecorder()

			DeleteUser(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var response map[string]string
			err := json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedMsg, response["message"])
		})
	}
}
