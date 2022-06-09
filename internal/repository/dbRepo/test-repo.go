package dbRepo

import (
	"database/sql"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/repository"
	"time"
)

type TestPostgresDBRepo struct {
	AppConfig *config.AppConfig
	DB        *sql.DB
}

func NewTestPostgresDBRepo(appConfigAsParam *config.AppConfig) repository.DatabaseRepository {
	return &TestPostgresDBRepo{
		AppConfig: appConfigAsParam,
	}
}

func (testPostgresDBRepo *TestPostgresDBRepo) AllUsers() bool {
	return true
}

func (testPostgresDBRepo *TestPostgresDBRepo) InsertReservation(reservation models.Reservation) (int, error) {
	var newReservationId int
	return newReservationId, nil
}

func (testPostgresDBRepo *TestPostgresDBRepo) InsertRoomRestriction(restriction models.RoomRestriction) error {
	return nil
}

func (testPostgresDBRepo *TestPostgresDBRepo) IsRoomAvailable(roomId int, startDate, endDate time.Time) (bool, error) {
	return false, nil
}

func (testPostgresDBRepo *TestPostgresDBRepo) AllAvailableRooms(startDate, endDate time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (testPostgresDBRepo *TestPostgresDBRepo) GetRoomById(roomId int) (models.Room, error) {
	var room models.Room
	return room, nil
}
