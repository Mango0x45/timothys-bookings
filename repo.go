package main

import (
	"database/sql"
)

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
		"User_Name" TEXT NOT NULL,
		"Booking_ID" INTEGER NOT NULL UNIQUE,
		"Room_ID" INTEGER NOT NULL UNIQUE,
		"Start_time" INTEGER NOT NULL,
		PRIMARY KEY("Booking_ID" AUTOINCREMENT),
		FOREIGN KEY("Room_ID") REFERENCES "Room" ("Room_ID")
	);
	`

	unwrap(r.db.Exec(query2))

	query3 := `
	INSERT INTO Room (Room_name) VALUES ('Oliver');
	INSERT INTO Room (Room_name) VALUES ('Amelia');
	INSERT INTO Room (Room_name) VALUES ('Ethan');
	INSERT INTO Room (Room_name) VALUES ('Sophia');
	INSERT INTO Room (Room_name) VALUES ('Liam');
	INSERT INTO Room (Room_name) VALUES ('Isabella');
	INSERT INTO Room (Room_name) VALUES ('Noah');
	INSERT INTO Room (Room_name) VALUES ('Mia');
	INSERT INTO Room (Room_name) VALUES ('Lucas');
	INSERT INTO Room (Room_name) VALUES ('Charlotte');
	`

	unwrap(r.db.Exec(query3))
}

type Booking struct {
	ID        int
	BookName  int
	RoomID    int
	StartTime string
}

func (r *SQLiteRepository) RegisterBooking(booking Booking) error {
	query := `
		INSERT INTO Booking (user_name, room_id, start_time)
		VALUES (?, ?, ?)
	`
	_, err := r.db.Exec(query, booking.BookName,
		booking.RoomID, booking.StartTime)
	return err
}

func (r *SQLiteRepository) GetAllBookings() ([]Booking, error) {
	query := `
		SELECT *
		FROM Booking
	`
	rows := unwrap(r.db.Query(query))

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		try(rows.Scan(&booking.ID, &booking.BookName,
			&booking.RoomID, &booking.StartTime))
		bookings = append(bookings, booking)
	}
	try(rows.Err())
	return bookings, nil
}
