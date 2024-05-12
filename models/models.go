package models

import "time"

type Trade struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	TraderId  string    `json:"traderId"`
	Symbol    string    `json:"symbol"`
	Quantity  float64   `json:"quantity"`
}

type UserTradeStat struct {
	TraderId    string  `json:"traderId"`
	TotalVolume float64 `json:"totalVolume"`
	Rank        int64   `json:"rank"`
}

type Symbol struct {
	Name      string `yaml:"name" json:"name"`
	Id        string `yaml:"id" json:"id"`
	Symbol    string `yaml:"symbol" json:"symbol"`
	ImagePath string `yaml:"imagePath" json:"imagePath"`
}

type Event struct {
	ID        string      `json:"id"`
	Timestamp time.Time   `json:"timestamp"`
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
}
