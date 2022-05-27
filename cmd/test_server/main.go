package main

import (
	"context"
	"fmt"

	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/test_server/internal/domain/event"

	"github.com/test_server/internal/infra/http"
	"github.com/test_server/internal/infra/http/controllers"
)

var settings = postgresql.ConnectionURL{
	Database: `test_db`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `1111`,
}

type Book struct {
	ID       uint   `db:"id,omitempty"`
	Title    string `db:"title"`
	AuthorID uint   `db:"author_id"`
}

type Author struct {
	ID        uint   `db:"id,omitempty"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
}

type BookAuthorSubject struct {
	// The book_id column was added to prevent collisions with the other "id"
	// columns from Author and Subject.
	BookID uint `db:"book_id"`

	Book   `db:",inline"`
	Author `db:",inline"`
}

// @title                       Test Server
// @version                     0.1.0
// @description                 Test Server boilerplate
func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	// Event
	eventRepository := event.NewRepository()
	eventService := event.NewService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)

	// HTTP Server
	err := http.Server(
		ctx,
		http.Router(
			eventController,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
	retrieveRecord()
	insertRecord()

}

func init() {
	// Use Open to access the database.
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// retrieve record

func retrieveRecord() {
	var booksQ []Book

	for _, book := range booksQ {
		fmt.Printf("Book %d:\t%q\n", book.ID, book.Title)
	}
	fmt.Println("")
}

func insertRecord(title string, Id int) []Book {
	var eaPoe Author
	book := Book{
		Title:    "The Crow",
		AuthorID: eaPoe.ID,
	}

	res, err = sess.SQL().
		InsertInto("books").
		Values(book). // Or Columns(c1, c2, c2, ...).Values(v1, v2, v2, ...).
		Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	}
	if res != nil {
		id, _ := res.LastInsertId()
		fmt.Printf("New book id: %d\n", id)
	}
	return book
}

func deleteRecord(id int) []Book {

	var eaPoe Author
	book := Book{
		Title:    "The Crow",
		AuthorID: eaPoe.ID,
	}
	q := sess.SQL().
		DeleteFrom("books").
		Where("title", "The Crow")
	fmt.Printf("Compiled query: %v\n", q)

	_, err = q.Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox\n", err)
	}
	return book
}
