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
	GetRoomById(roomId int) (models.Room, error)
	GetUserByID(userId int) (models.User, error)
	UpdateUser(user models.User) error
	Authenticate(email, testPassword string) (int, string, error)
}
