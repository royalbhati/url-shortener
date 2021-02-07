package database

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
	"github.com/royalbhati/urlshortener/app/config"
)

// Open knows how to open a database connection based on the configuration.
func Open(cfg *config.Config) (*sqlx.DB, error) {
	sslMode := "require"
	if cfg.Database.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.Database.User, cfg.Database.Password),
		Host:     cfg.Database.Host,
		Path:     cfg.Database.Name,
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}
