package models

var Gamers = make([]Gamer, 10, 15)

type Gamer struct {
	Name    string  `json:"name"`
	Skill   float32 `json:"skill"`
	Latency float32 `json:"latency"`
}