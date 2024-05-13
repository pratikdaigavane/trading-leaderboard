package models

import "time"

// Trade represents a trade event
type Trade struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	TraderId  string    `json:"traderId"`
	Symbol    string    `json:"symbol"`
	Quantity  float64   `json:"quantity"`
}

// UserTradeStat represents the user trade statistics
type UserTradeStat struct {
	TraderId    string  `json:"traderId"`
	TotalVolume float64 `json:"totalVolume"`
	Rank        int64   `json:"rank"`
}

// Symbol represents a symbol
type Symbol struct {
	Name      string `yaml:"name" json:"name"`
	Id        string `yaml:"id" json:"id"`
	Symbol    string `yaml:"symbol" json:"symbol"`
	ImagePath string `yaml:"imagePath" json:"imagePath"`
}

// Event represents an event, which can be a trade event or any other event
// This is essentially used by the event bus system
type Event struct {
	ID        string      `json:"id"`
	Timestamp time.Time   `json:"timestamp"`
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
}
