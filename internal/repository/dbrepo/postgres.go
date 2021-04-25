package dbrepo

import (
	"context"
	"github.com/firmanali/book-reservation/internal/models"
	"time"
)

func (p *postgresDbRepo) AllUsers() bool {
	return true
}

func (p *postgresDbRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date, end_date,
			room_id, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	err := p.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now()).Scan(&newId)

	if err != nil {
		return 0,err
	}
	return newId, nil
}

func (p *postgresDbRepo) InsertRoomRestriction(r models.RoomRestrictions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id,
			created_at, updated_at, restriction_id) values 
			($1, $2, $3, $4, $5, $6, $7)`

	_, err := p.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomId,
		r.ReservationId,
		time.Now(),
		time.Now(),
		r.RestrictionId )
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresDbRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var numRows int

	query := `	select count(id)
				from room_restrictions
				where room_id = $1 and $2 < end_date and $3 > start_date`
	row := p.DB.QueryRowContext(ctx, query,
		roomId, start, end)
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

func (p *postgresDbRepo) SearchAvailabilityAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room
	query := `select
				r.id, r.room_name
			from
				rooms r
			where
				r.id not in (select rr.room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date)`

	rows, err := p.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err = rows.Scan(&room.Id, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}
	return rooms, nil
}

func (p *postgresDbRepo) GetRoomById(id int) (models.Room,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := `select id, room_name, created_at, updated_at from rooms where id = $1`
	row := p.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&room.Id, &room.RoomName, &room.CreatedAt, &room.CreatedAt)
	if err != nil {
		return room, err
	}
	return room, nil
}