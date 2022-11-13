package service

import (
	"dev11/internal/domain/event"
	"log"
	"time"
)

type EventRepo interface {
	GetAll() (map[uint32]event.Event, error)
	Create(event event.Event) error
	Update(event event.Event) error
	Delete(id uint32) error
}

type CalendarService struct {
	EventRepo EventRepo
}

func NewCalendarService(repository EventRepo) CalendarService {
	return CalendarService{repository}
}

func (c *CalendarService) CreateEvent(event event.Event) error {
	err := c.EventRepo.Create(event)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (c *CalendarService) UpdateEvent(event event.Event) error {
	err := c.EventRepo.Update(event)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (c *CalendarService) DeleteEvent(id uint32) error {
	err := c.EventRepo.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (c *CalendarService) GetEventsForDay() ([]event.Event, error) {
	events, err := c.EventRepo.GetAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	start := getStartOfCurrentDay().Add(time.Nanosecond * -1)
	end := start.AddDate(0, 0, 1).Add(time.Nanosecond * -1)
	betweenEvents, err := getBetween(events, start, end)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return betweenEvents, nil
}

func (c *CalendarService) GetEventsForWeek() ([]event.Event, error) {
	events, err := c.EventRepo.GetAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	start := getStartOfCurrentDay().Add(time.Nanosecond * -1)
	end := start.AddDate(0, 0, 1).AddDate(0, 0, 7)
	betweenEvents, err := getBetween(events, start, end)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return betweenEvents, nil
}

func (c *CalendarService) GetEventsForMonth() ([]event.Event, error) {
	events, err := c.EventRepo.GetAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	start := getStartOfCurrentDay().Add(time.Nanosecond * -1)
	end := start.AddDate(0, 0, 1).AddDate(0, 1, 0)
	betweenEvents, err := getBetween(events, start, end)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return betweenEvents, nil
}

func getStartOfCurrentDay() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
}

func getBetween(events map[uint32]event.Event, start time.Time, end time.Time) ([]event.Event, error) {
	result := make([]event.Event, 0)
	for _, e := range events {
		if inTimeSpan(start, end, time.Time(e.Date)) {
			result = append(result, e)
		}

	}
	return result, nil
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
