package main

import "database/sql"

type SQLiteRepository struct {
	db *sql.DB
}

func (r *SQLiteRepository) migrate() error {
	query := `
	CREATE TABLE bookings
	(
		id INTEGER
		primary key AUTOINCREMENT,
		book_name  TEXT     NOT NULL,
		capacity   INTEGER  NOT NULL,
		room_id    INTEGER  NOT NULL,
		start_time DATETIME NOT NULL
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
	_, err := r.db.Exec(query, booking.BookName, booking.Capacity,
		booking.RoomID, booking.StartTime)
	return err
}

func (r *SQLiteRepository) GetAllBookings() ([]Booking, error) {
	query := `
		SELECT id, book_name, capacity, room_id, start_time
		FROM bookings
	`
	rows := unwrap(r.db.Query(query))

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		try(rows.Scan(&booking.ID, &booking.BookName, &booking.Capacity,
			&booking.RoomID, &booking.StartTime))
		bookings = append(bookings, booking)
	}
	try(rows.Err())
	return bookings, nil
}
