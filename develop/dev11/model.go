package main

import "time"

type event struct {
	userId      int       `json:"user_id"`
	userName    string    `json:"user_name"`
	discription string    `json:"discription"`
	date        time.Time `json:"date"`
}
