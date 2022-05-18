package models

// ReservationPageInputtedData holds personal user's data, which user gives on Reservation Page. This data are used for
// re-rendering, if user gives bad data
type ReservationPageInputtedData struct {
	FirstName string
	LastName  string
	Email     string `valid:"email"`
	Phone     string
}
