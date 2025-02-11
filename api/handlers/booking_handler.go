package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/saikumar-neelam/glofox_studio/internal/processors"
	"github.com/saikumar-neelam/glofox_studio/internal/structs"
	"github.com/saikumar-neelam/glofox_studio/internal/utils"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

type BookingRequest struct {
	ClassName  string `json:"class_name" validate:"required"`
	MemberName string `json:"member_name" validate:"required"`
	ClassDate  string `json:"class_date" validate:"required,dateformat"`
}

// validateDateFormat checks if the date is in the format YYYY-MM-DD
func validateDateFormat(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func init() {
	// Initialize the validator
	//we use validator to validate the input request after unmarshalling
	validate = validator.New()

	// Register custom validation for the date format
	//this helps in perfoming validation on datetime w.r.t format
	validate.RegisterValidation("dateformat", validateDateFormat)
}

func SendErrorResponse(w http.ResponseWriter, message string, details string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := structs.ErrorResponse{
		Error:   message,
		Details: details,
		Status:  statusCode,
	}

	utils.ErrorLogger.Println(errorResponse)
	// Encode the error response as JSON and write it to the response writer
	json.NewEncoder(w).Encode(errorResponse)
}

// BookClassHandler handles booking a class for a specific date
func BookClassHandler(w http.ResponseWriter, r *http.Request) {
	var request BookingRequest

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

	// Parse the class date
	classDate, err := time.Parse(DATEFORMAT, request.ClassDate)
	if err != nil {
		SendErrorResponse(w, "Invalid Data", "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	//check whether classDate provided is not past date
	if classDate.Before(time.Now().Truncate(24 * time.Hour)) {
		SendErrorResponse(w, "Invalid Data", "Invalid date. booking cannot be less than today", http.StatusBadRequest)
		return
	}

	// Call the booking service to create a booking
	booking, err := processors.BookClass(strings.ToLower(request.ClassName), request.MemberName, classDate)
	if err != nil {
		SendErrorResponse(w, "Unable to Process Request", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.InfoLogger.Printf("Booking for class %s confirmed for user %s on %s", request.ClassName, request.MemberName, classDate)

	// Return the created booking as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)

}

// GetBookingsByDateHandler handles fetching bookings for a specific class date
func GetBookingsByDateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classDateStr := vars["classDate"]

	// Parse the class date
	classDate, err := time.Parse(DATEFORMAT, classDateStr)
	if err != nil {
		SendErrorResponse(w, "Invalid date format. Use YYYY-MM-DD", err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service to fetch bookings
	bookings, err := processors.GetBookingsByDate(classDate)
	if err != nil {
		SendErrorResponse(w, "", err.Error(), http.StatusNotFound)
		return
	}

	// Return the bookings as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)
}
