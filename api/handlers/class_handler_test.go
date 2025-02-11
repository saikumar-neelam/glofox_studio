package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// Test for valid request
func TestCreateClassHandler_ValidRequest(t *testing.T) {

	requestBody := map[string]interface{}{
		"class_name": "Yoga",
		"start_date": "2025-02-20",
		"end_date":   "2025-03-15",
		"capacity":   10,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

// Test for missing required fields
func TestCreateClassHandler_MissingFields(t *testing.T) {
	requestBody := map[string]interface{}{
		"class_name": "",
		"start_date": "2025-02-15",
		"end_date":   "",
		"capacity":   10,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)

	if !strings.Contains(errorResponse.Details, "Name is missing or invalid") {
		t.Errorf("Expected 'Name is missing or invalid' error, got %v", response.Body.String())
	}

}

// Test for invalid date format
func TestCreateClassHandler_InvalidDateFormat(t *testing.T) {
	requestBody := map[string]interface{}{
		"class_name": "Yoga",
		"start_date": "15-02-2025", // Invalid format
		"end_date":   "2025-03-15",
		"capacity":   10,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)

	if !strings.Contains(errorResponse.Details, "StartDate is missing or invalid") {
		t.Errorf("Expected 'StartDate is missing or invalid' error, got %v", response.Body.String())
	}
}

// Test for CreateClass service error (e.g., class already exists)
// Below case will work only if we run the entire suite.
func TestCreateClassHandler_CreateClassServiceError(t *testing.T) {
	requestBody := map[string]interface{}{
		"class_name": "Yoga",
		"start_date": "2025-02-15",
		"end_date":   "2025-03-15",
		"capacity":   10,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusConflict, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)
	if !strings.Contains(errorResponse.Details, "class date conflicts with existing class schedule") {
		t.Errorf("Expected 'class date conflicts with existing class schedule' error, got %v", response.Body.String())
	}
}

// Test for invalid JSON body
func TestCreateClassHandler_InvalidJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer([]byte("{jskdfnsdfdf")))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	expected := "Invalid request body"

	json.Unmarshal(response.Body.Bytes(), &errorResponse)

	if !strings.Contains(errorResponse.Error, expected) {
		t.Errorf("Expected body %v, but got %v", expected, errorResponse.Details)
	}
}

// Test for missing required fields
func TestCreateClassHandler_InvalidStartDate(t *testing.T) {
	requestBody := map[string]interface{}{
		"class_name": "Pilates",
		"start_date": "2025-02-25",
		"end_date":   "2025-02-20",
		"capacity":   10,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	json.Unmarshal(response.Body.Bytes(), &errorResponse)

	if !strings.Contains(errorResponse.Details, "startDate cannot be greater than endDate") {
		t.Errorf("Expected 'startDate cannot be greater than endDate' error, got %v", errorResponse.Details)
	}
}
