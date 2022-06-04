package repository

import (
	"github.com/anonymfrominternet/Hotel/internal/models"
	"time"
)

type DatabaseRepository interface {
	AllUsers() bool

	InsertReservation(reservation models.Reservation) (int, error)

	InsertRoomRestriction(restriction models.RoomRestriction) error

	IsRoomAvailable(roomId int, startDate, endDate time.Time) (bool, error)

	AllAvailableRooms(startDate, endDate time.Time) ([]models.Room, error)
}
