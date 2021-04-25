package repository

import (
	"github.com/firmanali/book-reservation/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestrictions) error
	SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error)
	SearchAvailabilityAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(id int) (models.Room,error)
}
