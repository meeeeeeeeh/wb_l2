package main

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

type calendar struct {
	mu     sync.Mutex
	events map[string]map[string]time.Time //все события, ключ - имя пользователя
}

func (c *calendar) deleteEvent(user, event string) error {
	//c.mu.Lock()
	_, ok := c.events[user]
	if !ok {
		err := errors.New("no such username")
		return err
	}
	_, ok = c.events[user][event]
	if !ok {
		err := errors.New("no such event")
		return err
	}

	delete(c.events[user], event)

	//c.mu.Unlock()

	return nil
}

func (c *calendar) addEvent(date, user, newEventName string) error {
	d, err := time.Parse("2006-01-02T15:04", date)
	if err != nil {
		return err
	}

	ok := dateCheck(d)
	if !ok {
		err := errors.New("inalid date input")
		return err
	}

	// со слайсом:
	//можно переделать на мапу(????)
	newEvent := event{eventName: newEventName, date: d}

	//c.mu.Lock()
	_, ok = c.events[user]
	if !ok {
		c.events[user] = []event{newEvent}
	} else {
		c.events[user] = append(c.events[user], newEvent)
	}
	//c.mu.Unlock()

	return nil
}

func dateCheck(date time.Time) bool {
	if date.Year() > 2100 || date.Year() < 2000 {
		return false
	}
	return true

}

func newCalendar() *calendar {
	return &calendar{
		events: make(map[string]map[string]time.Time),
	}
}

func (c *calendar) createEventHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("template/create_event.html")
	b, _ := io.ReadAll(f)
	w.Write(b)

	if r.Method == http.MethodPost {

		err := c.addEvent(r.FormValue("date"), r.FormValue("user_name"), r.FormValue("event_name"))
		if err != nil {
			log.Fatalln(err)
		}

		// fmt.Println(c.events)

		// d, e := c.events["user"]
		// if e {
		// 	fmt.Println("user date", d[0].date)
		// }

	}

}

func (c *calendar) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("template/delete_event.html")
	b, _ := io.ReadAll(f)
	w.Write(b)

	if r.Method == http.MethodPost {

		err := c.deleteEvent(r.FormValue("user_name"), r.FormValue("event_name"))
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}
