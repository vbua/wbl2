package memory

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)
import "dev11/internal/domain/event"

// EventRepo реализует интерфейс в домене
type EventRepo struct {
	events map[uint32]event.Event
	*sync.RWMutex
}

func NewEventRepo() EventRepo {
	events := make(map[uint32]event.Event)
	return EventRepo{events, &sync.RWMutex{}}
}

func (e EventRepo) GetAll() (map[uint32]event.Event, error) {
	e.RLock()
	defer e.RUnlock()
	return e.events, nil
}

func (e EventRepo) Create(event event.Event) error {
	e.Lock()
	u := uuid.New()
	event.ID = u.ID()
	e.events[event.ID] = event
	e.Unlock()
	return nil
}

func (e EventRepo) Update(event event.Event) error {
	e.Lock()
	e.events[event.ID] = event
	e.Unlock()
	return nil
}

func (e EventRepo) Delete(id uint32) error {
	e.RLock()
	_, exists := e.events[id]
	e.RUnlock()
	if !exists {
		return fmt.Errorf("id %v not found", id)
	}
	e.Lock()
	delete(e.events, id)
	e.Unlock()
	return nil
}
