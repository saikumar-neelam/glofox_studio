package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/saikumar-neelam/glofox_studio/internal/processors"
	"github.com/saikumar-neelam/glofox_studio/internal/structs"
	"github.com/saikumar-neelam/glofox_studio/internal/utils"

	"net/http"
	"time"

	"github.com/go-playground/validator"
)

const DATEFORMAT = "2006-01-02"

// CreateClassHandler handles the creation of a new class
func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var request structs.ClassRequest
	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		SendErrorResponse(w, "Invalid request body", err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request fields
	err = validate.Struct(request)
	if err != nil {
		// If validation fails, extract validation errors and return specific error messages
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			// Return a clear message indicating the missing field or invalid date format
			errorMessage := fmt.Sprintf("%s is missing or invalid", e.Field())
			SendErrorResponse(w, "Invalid Data", errorMessage, http.StatusBadRequest)
			return
		}
	}

	// Parse the start and end date
	startDate, err := time.Parse(DATEFORMAT, request.StartDate)
	if err != nil {
		SendErrorResponse(w, "Invalid startDate format", err.Error(), http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(DATEFORMAT, request.EndDate)
	if err != nil {
		SendErrorResponse(w, "Invalid endDate format", err.Error(), http.StatusBadRequest)
		return
	}

	//check whether startdate/enddate is past date or not
	if startDate.Before(time.Now()) || endDate.Before(time.Now()) {
		SendErrorResponse(w, "Invalid startDate/endDate", "dates cannot be past date", http.StatusBadRequest)
		return
	}

	//check whether startdate is before enddate or not
	if startDate.After(endDate) {
		SendErrorResponse(w, "Invalid startDate/endDate", "startDate cannot be greater than endDate", http.StatusBadRequest)
		return
	}

	// Call the CreateClass processor to create class
	newClass, err := processors.CreateClass(strings.ToLower(request.ClassName), startDate, endDate, request.Capacity)
	if err != nil {
		SendErrorResponse(w, "Invalid Data", err.Error(), http.StatusConflict)
		return
	}

	utils.InfoLogger.Printf("Successfully created the class with classname %s from %s to %s", request.ClassName, startDate, endDate)

	// Return the created class in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClass)
}
