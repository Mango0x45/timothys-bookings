package main

import "database/sql"

type SQLiteRepository struct {
	db *sql.DB
}

func (r *SQLiteRepository) migrate() {
	query1 := `
	CREATE TABLE "Room" (
		"Room_ID" INTEGER NOT NULL UNIQUE,
		"Room_name" TEXT NOT NULL,
		PRIMARY KEY("Room_ID" AUTOINCREMENT)
 	);
	`
	unwrap(r.db.Exec(query1))

	query2 := `
	CREATE TABLE "Booking"
	(
		"User_Name" TEXT NOT NULL
		"Booking_ID" INTEGER NOT NULL UNIQUE,
		"Room_ID" INTEGER NOT NULL UNIQUE,
		"Start_time" INTEGER NOT NULL,
		PRIMARY KEY("Booking_ID" AUTOINCREMENT),
		FOREIGN KEY("Room_ID") REFERENCES "Room "("Room_ID"),
	);
	`

	unwrap(r.db.Exec(query2))
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
