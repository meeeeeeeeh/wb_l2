package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type event struct {
	//userName  string    `json:"user_name"`
	eventName string    `json:"event_name"`
	date      time.Time `json:"date"`
}

type calendar struct {
	mu     sync.Mutex
	events map[string][]event //все события, ключ - имя пользователя
}

func (c *calendar) deleteEvent(user, event string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.events) == 0 || len(c.events[user]) == 0 {
		err := errors.New("events db is empty")
		return err
	}

	_, ok := c.events[user]
	if !ok {
		err := errors.New("no such username")
		return err
	}

	// ищем имя ивента
	//если найдено сразу выходим и эррор - нил
	for idx, val := range c.events[user] {
		if val.eventName == event {
			//удаление ивента из слайса
			c.events[user] = append(c.events[user][:idx], c.events[user][idx+1:]...)
			return nil
		}
	}

	//если не найдено
	err := errors.New("no such event")
	return err

}

func (c *calendar) addEvent(date, user, newEventName string) error {
	d, err := time.Parse("2006-01-02T15:04", date)
	if err != nil {
		return err
	}

	ok := dateCheck(d)
	if !ok {
		err := errors.New("invalid date input")
		return err
	}

	newEvent := event{eventName: newEventName, date: d}

	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok = c.events[user]
	if !ok {
		c.events[user] = []event{newEvent}
	} else {
		//проверка есть ли такие же ивенты
		for _, val := range c.events[user] {
			if val == newEvent {
				err := errors.New("such event have been already created")
				return err
			}
		}

		c.events[user] = append(c.events[user], newEvent)
	}

	return nil
}

func (c *calendar) updateEvent(date, user, eventName string) error {
	d, err := time.Parse("2006-01-02T15:04", date)
	if err != nil {
		return err
	}

	ok := dateCheck(d)
	if !ok {
		err := errors.New("invalid date input")
		return err
	}

	updatedEvent := event{eventName: eventName, date: d}

	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok = c.events[user]
	if !ok {
		err := errors.New("no such username")
		return err
	}

	for idx, val := range c.events[user] {
		if val.eventName == eventName {
			c.events[user][idx] = updatedEvent
			return nil
		}
	}

	//если не найдено
	err = errors.New("no such event")
	return err

}

func dateCheck(date time.Time) bool {
	if date.Year() > 2100 || date.Year() < 2000 {
		return false
	}
	return true

}

func NewCalendar() *calendar {
	return &calendar{
		events: make(map[string][]event),
	}
}

func (c *calendar) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("template/create_event.html")
	b, _ := io.ReadAll(f)
	w.Write(b)

	if r.Method == http.MethodPost {

		err := c.addEvent(r.FormValue("date"), r.FormValue("user_name"), r.FormValue("event_name"))
		if err != nil {
			log.Fatalln(err)
		}

	}
	fmt.Println(c.events)

}

func (c *calendar) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("template/delete_event.html")
	b, _ := io.ReadAll(f)
	w.Write(b)

	if r.Method == http.MethodPost {

		err := c.deleteEvent(r.FormValue("user_name"), r.FormValue("event_name"))
		if err != nil {
			log.Fatalln(err)
		}

	}

	fmt.Println(c.events)
}

func (c *calendar) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("template/update_event.html")
	b, _ := io.ReadAll(f)
	w.Write(b)

	if r.Method == http.MethodPost {

		err := c.updateEvent(r.FormValue("date"), r.FormValue("user_name"), r.FormValue("event_name"))
		if err != nil {
			log.Fatalln(err)
		}

	}
	fmt.Println(c.events)
}

func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("method: ", r.Method, "path: ", r.URL.EscapedPath())

	}
}
