package handlers

import (
	"dev11/internal/models"
	"dev11/internal/services"
	"encoding/json"
	"net/http"
)

type EventHandler struct {
	eventService services.EventService
}

func NewEventHandler(eventService services.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

// Реализация методов обработчиков HTTP-запросов...

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Парсим входные данные в ивент
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обращаемся к сервису
	err = h.eventService.CreateEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа с данными в формате JSON
	w.Header().Set("Content-Type", "application/json")
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// Парсим входные данные в ивент
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обращаемся к сервису
	err = h.eventService.UpdateEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа с данными в формате JSON
	w.Header().Set("Content-Type", "application/json")

}

func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	// Получение данных о событиях
	data, err := h.eventService.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Сериализация данных в формат JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправка ответа с данными в формате JSON
	w.Write(jsonData)
}
func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	h.eventService.DeleteEvent()
	//
	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")
}
