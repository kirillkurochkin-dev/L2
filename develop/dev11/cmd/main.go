package main

import (
	"database/sql"
	"dev11/internal/handlers"
	"dev11/internal/services/impl"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	fmt.Println("ЗАПУСК")
	// Подключение к БД

	connStr := "user=postgres dbname=calendar password=postgres sslmode=disable"

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Сервис и контроллер
	eventService := impl.NewPostgreSQLService(db)
	eventHandler := handlers.NewEventHandler(eventService)

	// Обработка

	// POST
	http.HandleFunc("/create_event", eventHandler.CreateEvent)
	// POST
	//http.HandleFunc("/update_event", ...)
	// POST
	//http.HandleFunc("/delete_event", ...)
	// GET
	http.HandleFunc("/events", eventHandler.GetAllEvents)
	// GET
	//http.HandleFunc("/events_for_day", ...)
	// GET
	//http.HandleFunc("/events_for_week", ...)
	// GET
	//http.HandleFunc("/events_for_month", ...)

	// Слушаем порт
	http.ListenAndServe(":8080", nil)

}
