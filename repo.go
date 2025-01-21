package main

import "database/sql"

type SQLiteRepository struct {
	db *sql.DB
}

func (r *SQLiteRepository) migrate() error {
	query := `
	create table bookings
	(
    id INTEGER
    primary key autoincrement,
    book_name  TEXT     not null,
    capacity   INTEGER  not null,
    room_id    INTEGER  not null,
    start_time DATETIME not null
	);
    `

	_, err := r.db.Exec(query)
	return err
}

type Booking struct {
	ID        int
	BookName  string
	Capacity  int
	RoomID    int
	StartTime string
}

func (r *SQLiteRepository) RegisterBooking(booking Booking) error {
	query := `
	INSERT INTO bookings (book_name, capacity, room_id, start_time)
	VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, booking.BookName, booking.Capacity, booking.RoomID, booking.StartTime)
	return err
}

func (r *SQLiteRepository) GetAllBookings() ([]Booking, error) {
	query := `
	SELECT id, book_name, capacity, room_id, start_time
	FROM bookings
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.ID, &booking.BookName, &booking.Capacity, &booking.RoomID, &booking.StartTime)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil

}
