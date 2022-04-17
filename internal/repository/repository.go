package repository

import "github.com/anonymfrominternet/Hotel/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(reservation models.Reservation) error
}
