package controller

// func TestGetDataUser(t *testing.T) {
// 	// Mock request
// 	req := httptest.NewRequest(http.MethodGet, "/getdatauser", nil)
// 	req.Header.Set("Authorization", "Bearer dummyToken")
// 	resp := httptest.NewRecorder()

// 	// Mock fungsi Decode dan GetOneDoc
// 	mockDecode := func(publicKey, token string) (Payload, error) {
// 		return Payload{Id: "123456789", Alias: "Test User"}, nil
// 	}
// 	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
// 		return model.Userdomyikado{PhoneNumber: "123456789", Name: "Test User"}, nil
// 	}

// 	// Ganti fungsi asli dengan mock
// 	originalDecode := watoken.Decode
// 	originalGetOneDoc := atdb.GetOneDoc
// 	defer func() { watoken.Decode = originalDecode; atdb.GetOneDoc = originalGetOneDoc }()
// 	watoken.Decode = mockDecode
// 	atdb.GetOneDoc = mockGetOneDoc

// 	// Panggil fungsi
// 	GetDataUser(resp, req)

// 	// Validasi hasil
// 	if resp.Code != http.StatusOK {
// 		t.Errorf("expected status OK, got %d", resp.Code)
// 	}
// }

// func TestPostDataUser(t *testing.T) {
// 	// Mock request body
// 	user := model.Userdomyikado{
// 		Name:        "Test User",
// 		PhoneNumber: "123456789",
// 		Email:       "test@example.com",
// 	}
// 	body, _ := json.Marshal(user)
// 	req := httptest.NewRequest(http.MethodPost, "/postdatauser", bytes.NewReader(body))
// 	req.Header.Set("Authorization", "Bearer dummyToken")
// 	resp := httptest.NewRecorder()

// 	// Mock fungsi Decode, GetOneDoc, dan InsertOneDoc
// 	mockDecode := func(publicKey, token string) (Payload, error) {
// 		return Payload{Id: "123456789", Alias: "Test User"}, nil
// 	}
// 	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.Userdomyikado, error) {
// 		return model.Userdomyikado{}, fmt.Errorf("not found")
// 	}
// 	mockInsertOneDoc := func(dbConn interface{}, collection string, doc interface{}) (string, error) {
// 		return "mocked_id", nil
// 	}

// 	originalDecode := watoken.Decode
// 	originalGetOneDoc := atdb.GetOneDoc
// 	originalInsertOneDoc := atdb.InsertOneDoc
// 	defer func() {
// 		watoken.Decode = originalDecode
// 		atdb.GetOneDoc = originalGetOneDoc
// 		atdb.InsertOneDoc = originalInsertOneDoc
// 	}()
// 	watoken.Decode = mockDecode
// 	atdb.GetOneDoc = mockGetOneDoc
// 	atdb.InsertOneDoc = mockInsertOneDoc

// 	// Panggil fungsi
// 	PostDataUser(resp, req)

// 	// Validasi hasil
// 	if resp.Code != http.StatusOK {
// 		t.Errorf("expected status OK, got %d", resp.Code)
// 	}
// }

// func TestDeleteQRHistory(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodDelete, "/deleteqrhistory?id=mocked_id", nil)
// 	req.Header.Set("Authorization", "Bearer dummyToken")
// 	resp := httptest.NewRecorder()

// 	// Mock fungsi
// 	mockDecode := func(publicKey, token string) (Payload, error) {
// 		return Payload{Id: "123456789", Alias: "Test User"}, nil
// 	}
// 	mockGetOneDoc := func(dbConn interface{}, collection string, filter interface{}) (model.QrHistory, error) {
// 		return model.QrHistory{ID: primitive.NewObjectID(), Name: "ExistingQR"}, nil
// 	}
// 	mockDeleteOneDoc := func(dbConn interface{}, collection string, filter interface{}) (int64, error) {
// 		return 1, nil
// 	}

// 	originalDecode := watoken.Decode
// 	originalGetOneDoc := atdb.GetOneDoc
// 	originalDeleteOneDoc := atdb.DeleteOneDoc
// 	defer func() {
// 		watoken.Decode = originalDecode
// 		atdb.GetOneDoc = originalGetOneDoc
// 		atdb.DeleteOneDoc = originalDeleteOneDoc
// 	}()
// 	watoken.Decode = mockDecode
// 	atdb.GetOneDoc = mockGetOneDoc
// 	atdb.DeleteOneDoc = mockDeleteOneDoc

// 	// Panggil fungsi
// 	DeleteQRHistory(resp, req)

// 	// Validasi hasil
// 	if resp.Code != http.StatusOK {
// 		t.Errorf("expected status OK, got %d", resp.Code)
// 	}
// }