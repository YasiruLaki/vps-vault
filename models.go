package main

import "time"

type VPSData struct {
	Name string `json:"name"`
	IP string `json:"ip"`
	User string `json:"username"`
	Port int `json:"port"`
	CreatedAt time.Time `json:"created_at"`
}
