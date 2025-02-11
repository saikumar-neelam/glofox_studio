package processors

import (
	"testing"
	"time"
)

func TestCreateClass(t *testing.T) {
	className := "Sai Kumar"
	startDate, _ := time.Parse(DATEFORMAT, "2025-02-20")
	endDate, _ := time.Parse(DATEFORMAT, "2025-02-28")
	capacity := 100
	// Create class
	class, err := CreateClass(className, startDate, endDate, capacity)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if class.ClassName != className {
		t.Fatalf("expected member name %s, got %s", className, class.ClassName)
	}

	if !class.StartDate.Equal(startDate) {
		t.Fatalf("expected class date %v, got %v", startDate, class.StartDate)
	}
	if !class.EndDate.Equal(endDate) {
		t.Fatalf("expected class date %v, got %v", endDate, class.EndDate)
	}

}
