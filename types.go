package main

import "time"

type Language struct {
	ID   int
	Name string
}

type CodeReport struct {
	ID          int
	Request     int
	Language_id int
	Score       int
	Percentage  float64
	Created_At  time.Time
}