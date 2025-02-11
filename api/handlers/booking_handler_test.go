package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/saikumar-neelam/glofox_studio/internal/structs"

	"github.com/gorilla/mux"
)

var errorResponse structs.ErrorResponse

// executeRequest will create a mux router to perform the test cases
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	// Route to create a new class
	r.HandleFunc("/classes", CreateClassHandler).Methods(http.MethodPost)
	r.HandleFunc("/bookings", BookClassHandler).Methods(http.MethodPost)
	r.HandleFunc("/bookings/{classDate}", GetBookingsByDateHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	return rr
}

// function checkResponseCode used to verifiy the expected and actual status codes
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d", expected, actual)
	}
}

func TestBookClass_Success(t *testing.T) {
	payload := `{"member_name":"Sai Kumar", "class_date":"2025-02-15", "class_name": "Yoga"}`
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)

	if response.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, response.Code)
	}
}

func TestBookClass_InvalidJSON(t *testing.T) {
	payload := `{"member_name":"Sai Kumar", "class_date":"2025-02-15"`
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestBookClass_MissingMemberName(t *testing.T) {

	payload := `{"class_date":"2025-02-15", "class_name": "Yoga"}`
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)
	if !strings.Contains(errorResponse.Details, "MemberName is missing or invalid") {
		t.Errorf("Expected 'MemberName is missing or invalid' error, got %v", response.Body.String())
	}
}

func TestBookClass_InvalidDateFormat(t *testing.T) {
	payload := `{"member_name":"Sai Kumar", "class_date":"15-02-2025", "class_name": "Yoga"}`
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)
	if !strings.Contains(errorResponse.Details, "ClassDate is missing or invalid") {
		t.Errorf("Expected 'ClassDate is missing or invalid' error, got %v", errorResponse.Details)
	}
}

func TestBookClass_InvalidClassDate(t *testing.T) {
	payload := `{"member_name":"Sai Kumar", "class_date":"2025-02-01", "class_name": "Yoga"}`
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
	json.Unmarshal(response.Body.Bytes(), &errorResponse)
	if !strings.Contains(errorResponse.Details, "Invalid date. booking cannot be less than today") {
		t.Errorf("Expected 'Invalid date. booking cannot be less than today' error, got %v", errorResponse.Details)
	}
}

func TestBookClass_EmptyMemberName(t *testing.T) {
	payload := `{"member_name":"", "class_date":"2025-02-15", "class_name": "Yoga"}`
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
	json.Unmarshal(response.Body.Bytes(), &errorResponse)
	if !strings.Contains(errorResponse.Details, "MemberName is missing or invalid") {
		t.Errorf("Expected 'MemberName is missing or invalid' error, got %v", errorResponse.Details)
	}
}
