package shortener

import (
	"time"
)

type Info struct {
	ID          string    `db:"id" json:"id"`
	ShortUrl    string    `db:"short_url" json:"short_url"`
	LongUrl     string    `db:"long_url" json:"long_url"`
	DateCreated time.Time `db:"date_created" json:"created_at"`
	DateUpdated time.Time `db:"date_updated" json:"updated_at"`
}

type NewItem struct {
	LongUrl  string `db:"long_url" json:"long_url"`
	ShortUrl string `db:"short_url" json:"short_url"`
}
