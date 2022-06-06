package dbRepo

import (
	context2 "context"
	"database/sql"
	"fmt"
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

// InsertReservation inserts new items into the reservations table
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

// InsertRoomRestriction inserts new items in the room_restriction table
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

// IsRoomAvailable checks if a room is available for a user in period of user's dates
func (postgresDBRepo *postgresDBRepo) IsRoomAvailable(roomId int, startDate, endDate time.Time) (bool, error) {
	context, cancel := context2.WithTimeout(context2.Background(), 3*time.Second)
	defer cancel()

	var numRows int
	query := `select
					count(id)
				from room_restrictions
				where 
				    room_id = $1
				    2$ < start_date or $3 > end_date;
			`
	row := postgresDBRepo.DB.QueryRowContext(context, query,
		roomId,
		endDate,
		startDate,
	)

	fmt.Println("number of rows is ", numRows)

	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows > 0 {
		return true, nil
	}

	return false, nil
}

// AllAvailableRooms gets all available rooms, which correspond to given dates
func (postgresDBRepo *postgresDBRepo) AllAvailableRooms(startDate, endDate time.Time) ([]models.Room, error) {
	context, cancel := context2.WithTimeout(context2.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
			select
				r.id, r.room_name
			from
				rooms r
			where r.id in (select rr.room_id from room_restrictions rr where $1 > rr.end_date or $2 < rr.start_date);
`
	rows, err := postgresDBRepo.DB.QueryContext(context, query, startDate, endDate)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, nil
		}

		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomById gets a room by given id from db
func (postgresDBRepo *postgresDBRepo) GetRoomById(roomId int) (models.Room, error) {
	context, cancel := context2.WithTimeout(context2.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := `
			select id, room_name, created_at, updated_at from rooms where id = $1
`
	row := postgresDBRepo.DB.QueryRowContext(context, query,
		roomId)

	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return room, err
	}
	return room, nil

}

func NewPostgresDBRepo(appConfigAsParam *config.AppConfig, db *sql.DB) repository.DatabaseRepository {
	return &postgresDBRepo{
		AppConfig: appConfigAsParam,
		DB:        db,
	}
}
