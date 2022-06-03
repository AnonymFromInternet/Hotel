package repository

import "github.com/anonymfrominternet/Hotel/internal/models"

type DatabaseRepository interface {
	AllUsers() bool

	InsertReservation(reservation models.Reservation) (int, error)

	InsertRoomRestriction(restriction models.RoomRestriction) error
}
