package routers

import (
	"net/http"

	"github.com/saikumar-neelam/glofox_studio/api/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter sets up the API routes using gorilla/mux
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Route to create a new class
	r.HandleFunc("/classes", handlers.CreateClassHandler).Methods(http.MethodPost)

	//Route to book a class
	r.HandleFunc("/bookings", handlers.BookClassHandler).Methods(http.MethodPost)

	//Route to get the number of bookings of different classes on specific date
	r.HandleFunc("/bookings/{classDate}", handlers.GetBookingsByDateHandler).Methods("GET")
	return r
}
