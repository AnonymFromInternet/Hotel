package dbrepo

import (
	"github.com/anonymfrominternet/Hotel/internal/models"
	"golang.org/x/net/context"
	"time"
)

// These methods are implementation methods for postgresDBRepo struct

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(reservation models.Reservation) error {

	context, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, 
	updated_at)   values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	_, err := m.DB.ExecContext(
		context,
		statement,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
