package processors

import (
	"errors"
	"sync"
	"time"

	"github.com/saikumar-neelam/glofox_studio/internal/structs"
)

const DATEFORMAT = "2006-01-02"

var DateWiseoverallBookings = make(map[string]map[string][]structs.Booking)

var bookingsMutex sync.Mutex

// bookclass is a function which implements booking a class for a member
// input name and class date
// output booking struct, error
func BookClass(class_name, member_name string, classDate time.Time) (structs.Booking, error) {

	defer bookingsMutex.Unlock()
	bookingsMutex.Lock()
	newBooking := structs.Booking{MemberName: member_name, ClassDate: classDate, ClassName: class_name}

	date := classDate.Format(DATEFORMAT)

	//check already date wise any bookings are there
	//if no bookings found for the date then initialize it
	if _, ok := DateWiseoverallBookings[date]; !ok {
		DateWiseoverallBookings[date] = make(map[string][]structs.Booking)
	}

	//check if class not found or not on the date.
	if _, ok := DateWiseoverallBookings[date][class_name]; !ok {
		//looping through classes to find any classname matches with the start and end date range
		for _, existingClass := range classes {
			if existingClass.ClassName == class_name && ((classDate.Equal(existingClass.StartDate) || classDate.After(existingClass.StartDate)) && (classDate.Equal(existingClass.EndDate) || classDate.Before(existingClass.EndDate))) {
				// Initialize the class bookings for the date if it does not exist
				DateWiseoverallBookings[date][class_name] = []structs.Booking{}
				DateWiseoverallBookings[date][class_name] = append(DateWiseoverallBookings[date][class_name], newBooking)
				return newBooking, nil
			}
		}
	}

	DateWiseoverallBookings[date][class_name] = append(DateWiseoverallBookings[date][class_name], newBooking)
	return newBooking, nil
}

// GetbookingsByDate function will return the total number of bookings
// done on particular date
// input classdate
// output list of bookings
func GetBookingsByDate(classDate time.Time) (map[string][]structs.Booking, error) {

	date := classDate.Format(DATEFORMAT)
	//check whether anybookings are there
	if len(DateWiseoverallBookings[date]) == 0 {
		return nil, errors.New("no bookings available for the selected date")
	}
	return DateWiseoverallBookings[date], nil
}
