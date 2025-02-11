package processors

import (
	"testing"
	"time"
)

func TestBookClass(t *testing.T) {
	memberName := "Sai Kumar"
	ClassName := "Yoga"
	classDate, _ := time.Parse("2006-01-02", "2025-02-22")

	// Create booking
	booking, err := BookClass(ClassName, memberName, classDate)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if booking.MemberName != memberName {
		t.Fatalf("expected member name %s, got %s", memberName, booking.MemberName)
	}

	if !booking.ClassDate.Equal(classDate) {
		t.Fatalf("expected class date %v, got %v", classDate, booking.ClassDate)
	}
}
