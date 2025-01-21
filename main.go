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

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("POST /book", createbooking)
	http.HandleFunc("GET /", root)
	http.HandleFunc("GET /book", booking)

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
	db := unwrap(sql.Open("sqlite3", ":memory:"))
	m := SQLiteRepository{db}
	m.migrate()
	return m
}

func booking(w http.ResponseWriter, r *http.Request) {
	// c, err := r.Cookie("user")
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, http.ErrNoCookie):
	// 		http.Error(w, "Cookie not found", http.StatusBadRequest)
	// 	}
	// }
	try(templates["book"].Execute(w, nil))
}