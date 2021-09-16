package impl

import (
	"database/sql"
	"dev11/internal/models"
	"time"
)

type PostgreSQLService struct {
	DB *sql.DB
}

func NewPostgreSQLService(db *sql.DB) *PostgreSQLService {
	return &PostgreSQLService{DB: db}
}

// Реализация методов интерфейса EventService для PostgreSQLService

func (p PostgreSQLService) CreateEvent(event models.Event) error {
	query := "INSERT INTO events (user_id, title, start_time, end_time) VALUES ($1, $2, $3, $4)"
	_, err := p.DB.Exec(query, event.UserID, event.Title, event.StartTime, event.EndTime)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgreSQLService) UpdateEvent(event models.Event) error {
	query := "UPDATE events SET user_id = $1, title = $2, start_time = $3, end_time = $4 WHERE id = $5"
	_, err := p.DB.Exec(query, event.UserID, event.Title, event.StartTime, event.EndTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgreSQLService) DeleteEvent(id int) error {
	query := "DELETE FROM events WHERE id = $1"
	_, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgreSQLService) GetAllEvents() ([]*models.Event, error) {
	query := "SELECT * FROM events"
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Создание среза для хранения событий
	var events []*models.Event

	for rows.Next() {
		// Создание переменных для сканирования значений из строк
		var event models.Event
		err := rows.Scan(&event.ID, &event.UserID, &event.Title, &event.StartTime, &event.EndTime)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	// Проверка ошибок после завершения итерации
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Возвращение событий
	return events, nil
}

func (p PostgreSQLService) GetEventsForDay(userID int, date time.Time) ([]*models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgreSQLService) GetEventsForWeek(userID int, date time.Time) ([]*models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgreSQLService) GetEventsForMonth(userID int, date time.Time) ([]*models.Event, error) {
	//TODO implement me
	panic("implement me")
}
