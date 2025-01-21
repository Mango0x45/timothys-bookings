package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var templates = make(map[string]*template.Template)
var sqldb = setUpDatabase()

func main() {
	for _, e := range unwrap(os.ReadDir("./templates")) {
		base := strings.TrimSuffix(e.Name(), ".html")
		file := "templates/" + e.Name()
		templates[base] = template.Must(template.ParseFiles(file))
	}
	defer sqldb.db.Close()
	http.HandleFunc("POST /book", createbooking)
	http.HandleFunc("/", root)
	try(http.ListenAndServe(":6969", nil))

}

func try(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func unwrap[T any](x T, err error) T {
	try(err)
	return x
}

func root(w http.ResponseWriter, r *http.Request) {
	try(templates["index"].Execute(w, nil))
}

func createbooking(w http.ResponseWriter, r *http.Request) {
	x := r.Form.Get("")
	fmt.Println(x)
	booking := Booking{
		1,
		"NAME",
		2,
		1,
		"2025-01-25 14:30:00",
	}
	sqldb.RegisterBooking(booking)
	bookings := unwrap(sqldb.GetAllBookings())
	fmt.Fprintf(w, "%s book", bookings[0].BookName)

}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func setUpDatabase() SQLiteRepository {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(db)
	m := SQLiteRepository{db}
	m.migrate()
	return m
}
