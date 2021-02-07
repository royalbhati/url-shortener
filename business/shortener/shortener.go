package shortener

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const uniqueKey = "key"

type UrlShortener struct {
	log *log.Logger
	db  *sqlx.DB
	rdb *redis.Client
}

// New constructs a Product for api access.
func New(log *log.Logger, db *sqlx.DB, redis *redis.Client) UrlShortener {
	return UrlShortener{
		log: log,
		db:  db,
		rdb: redis,
	}
}

func (u UrlShortener) CreateShortURL(ctx context.Context, item NewItem) (Info, error) {
	shortURL := item.ShortUrl
	if item.ShortUrl == "" {
		id := GetUniqueKey(ctx, u.log, u.rdb)
		possibleChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

		shorturl := ""

		for id > 0 {
			shorturl += string(possibleChars[id%62])
			id = id / 62
		}
		shortURL = shorturl
	} else {
		val, err := u.GetBigURL(ctx, shortURL)

		if err != nil {
			return Info{}, errors.Wrap(err, "checking existing short url in db")
		}
		if val != "" {
			u.log.Println("", val, errors.Errorf("Given key Already exists"))
			return Info{}, errors.Errorf("Given key Already exists")
		}
	}
	u.log.Println("HEREEE")

	newI := Info{
		ID:          uuid.New().String(),
		ShortUrl:    shortURL,
		LongUrl:     item.LongUrl,
		DateCreated: time.Now().UTC(),
		DateUpdated: time.Now().UTC(),
	}

	const q = `
	INSERT INTO urls
		(id,short_url, long_url, date_created, date_updated)
	VALUES
		($1, $2, $3, $4, $5)`
	if _, err := u.db.ExecContext(ctx, q, newI.ID, newI.ShortUrl, newI.LongUrl, newI.DateCreated, newI.DateUpdated); err != nil {
		return Info{}, errors.Wrap(err, "inserting url")
	}
	u.log.Println("SHORT URL", shortURL)
	return newI, nil

}

func (u UrlShortener) GetBigURL(ctx context.Context, key string) (string, error) {
	const q = `
	SELECT
		*
	FROM
		urls
	WHERE
		short_url = $1`

	var item Info
	if err := u.db.GetContext(ctx, &item, q, key); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.Wrap(err, "Not present")
		}
		return "", errors.Wrap(err, "querying db")
	}
	u.log.Println("HEYY", item)

	return item.LongUrl, nil
}

func GetUniqueKey(ctx context.Context, log *log.Logger, rdb *redis.Client) int64 {
	val, err := rdb.Incr(ctx, uniqueKey).Result()
	if err != nil {
		log.Fatal("Unique ID doesnt exist ")
	}
	return val
}
