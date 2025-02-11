package structs

import "time"

// Class represents a studio class
type Class struct {
	ID        int       `json:"id"` //unique identifier which can be used when we store data in database
	ClassName string    `json:"class_name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

type Booking struct {
	MemberName string    `json:"member_name"`
	ClassDate  time.Time `json:"class_date"`
	ClassName  string    `json:"class_name"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"status"`
}

type ClassRequest struct {
	ClassName string `json:"class_name" validate:"required"`
	StartDate string `json:"start_date" validate:"required,dateformat"`
	EndDate   string `json:"end_date" validate:"required,dateformat"`
	Capacity  int    `json:"capacity" validate:"required"`
}
