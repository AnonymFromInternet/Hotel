package dbRepo

import (
	context2 "context"
	"database/sql"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/repository"
	"time"
)

type postgresDBRepo struct {
	AppConfig *config.AppConfig
	DB        *sql.DB
}

func (postgresDBRepo *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts new items into the Reservations table
func (postgresDBRepo *postgresDBRepo) InsertReservation(reservation models.Reservation) (int, error) {
	context, cancel := context2.WithTimeout(context2.Background(), 3*time.Second)
	defer cancel()

	var newReservationId int
	statement := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := postgresDBRepo.DB.QueryRowContext(context, statement,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomId,
		time.Now(),
		time.Now(),
	).Scan(&newReservationId)
	if err != nil {
		return 0, err
	}

	return newReservationId, nil
}

func (postgresDBRepo *postgresDBRepo) InsertRoomRestriction(restriction models.RoomRestriction) error {
	context, cancel := context2.WithTimeout(context2.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := postgresDBRepo.DB.ExecContext(context, statement,
		restriction.StartDate,
		restriction.EndDate,
		restriction.RoomId,
		restriction.ReservationId,
		restriction.RestrictionId,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func NewPostgresDBRepo(appConfigAsParam *config.AppConfig, db *sql.DB) repository.DatabaseRepository {
	return &postgresDBRepo{
		AppConfig: appConfigAsParam,
		DB:        db,
	}
}
