package server

import (
	"dev11/internal/domain/event"
	"net/http"
	"strconv"
	"time"
)

type CalendarService interface {
	CreateEvent(event event.Event) error
	UpdateEvent(event event.Event) error
	DeleteEvent(id uint32) error
	GetEventsForDay() ([]event.Event, error)
	GetEventsForWeek() ([]event.Event, error)
	GetEventsForMonth() ([]event.Event, error)
}

type Handler struct {
	CalendarService CalendarService
}

func NewHandler(service CalendarService) Handler {
	return Handler{service}
}

func (h *Handler) CreateRoutes() *http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventsForDay)
	mux.HandleFunc("/events_for_week", h.eventsForWeek)
	mux.HandleFunc("/events_for_month", h.eventsForMonth)
	handler := logging(mux)
	return &handler
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		sendResponse(true, "require method POST", http.StatusMethodNotAllowed, w)
		return
	}
	e := &event.Event{}
	formEventTitle := r.FormValue("title")
	if formEventTitle == "" {
		sendResponse(true, "No title", http.StatusBadRequest, w)
	}
	formEventDate := r.FormValue("date")
	date, err := time.Parse("2006-01-02", formEventDate)
	if err != nil {
		sendResponse(true, "Bad date", http.StatusBadRequest, w)
	}

	e.Title = formEventTitle
	e.Date = date
	err = h.CalendarService.CreateEvent(*e)
	if err != nil {
		sendResponse(true, err.Error(), http.StatusServiceUnavailable, w)
		return
	}
	sendResponse(false, "Event created successfully", http.StatusOK, w)
	return
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		sendResponse(true, "require method POST", http.StatusMethodNotAllowed, w)
		return
	}
	e := &event.Event{}
	formEventId := r.FormValue("id")
	eventId, err := strconv.Atoi(formEventId)
	if err != nil {
		sendResponse(true, "Invalid id", http.StatusBadRequest, w)
		return
	}
	formEventTitle := r.FormValue("title")
	if formEventTitle == "" {
		sendResponse(true, "No title", http.StatusBadRequest, w)
	}
	formEventDate := r.FormValue("date")
	date, err := time.Parse("2006-01-02", formEventDate)
	if err != nil {
		sendResponse(true, "Bad date", http.StatusBadRequest, w)
	}

	e.ID = uint32(eventId)
	e.Title = formEventTitle
	e.Date = date

	err = h.CalendarService.UpdateEvent(*e)
	if err != nil {
		sendResponse(true, err.Error(), http.StatusServiceUnavailable, w)
		return
	}
	sendResponse(false, "Event updated successfully", http.StatusOK, w)
	return
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		sendResponse(true, "require method POST", http.StatusMethodNotAllowed, w)
		return
	}
	formEventId := r.FormValue("id")
	eventId, err := strconv.Atoi(formEventId)
	if err != nil {
		sendResponse(true, "Invalid id", http.StatusBadRequest, w)
		return
	}

	err = h.CalendarService.DeleteEvent(uint32(eventId))
	if err != nil {
		sendResponse(true, err.Error(), http.StatusServiceUnavailable, w)
		return
	}
	sendResponse(false, "Event deleted successfully", http.StatusOK, w)
	return
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		sendResponse(true, "require method GET", http.StatusMethodNotAllowed, w)
		return
	}
	events, err := h.CalendarService.GetEventsForDay()
	if err != nil {
		sendResponse(true, err.Error(), http.StatusServiceUnavailable, w)
		return
	}
	sendResponse(false, events, http.StatusOK, w)
	return
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		sendResponse(true, "require method GET", http.StatusMethodNotAllowed, w)
		return
	}
	events, err := h.CalendarService.GetEventsForWeek()
	if err != nil {
		sendResponse(true, err.Error(), http.StatusServiceUnavailable, w)
		return
	}
	sendResponse(false, events, http.StatusOK, w)
	return
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		sendResponse(true, "require method GET", http.StatusMethodNotAllowed, w)
		return
	}
	events, err := h.CalendarService.GetEventsForMonth()
	if err != nil {
		sendResponse(true, err.Error(), http.StatusServiceUnavailable, w)
		return
	}
	sendResponse(false, events, http.StatusOK, w)
	return
}
