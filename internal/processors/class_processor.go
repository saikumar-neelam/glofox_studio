package processors

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/saikumar-neelam/glofox_studio/internal/structs"
)

var classes []structs.Class
var classID int64 = 1

// CreateClass adds a new class to the list
// input name, startDate, endDate, capacity
// output classobject, error
func CreateClass(name string, startDate, endDate time.Time, capacity int) (structs.Class, error) {

	// Before adding the new class, looping through the existing classes and
	// check if any dates are overlapping. If any conflict is found, an error message is returned,
	// and the class is not created.

	for _, existingClass := range classes {
		if existingClass.ClassName == name {
			if (startDate.Before(existingClass.EndDate) && endDate.After(existingClass.StartDate)) ||
				startDate.Equal(existingClass.StartDate) || endDate.Equal(existingClass.EndDate) {
				return structs.Class{}, errors.New("class date conflicts with existing class schedule")
			}
		}
	}

	newClass := structs.Class{
		ID:        int(classID),
		ClassName: name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}

	//atomic incrementation helps race condition while working with concurrent requests
	atomic.AddInt64(&classID, 1)

	classes = append(classes, newClass)
	return newClass, nil
}
