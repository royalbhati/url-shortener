package shortener

import (
	"github.com/dimiro1/darwin"
	"github.com/jmoiron/sqlx"
)

func Migrate(db *sqlx.DB) error {
	driver := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})
	d := darwin.New(driver, migrations, nil)
	return d.Migrate()
}

var migrations = []darwin.Migration{
	{
		Version:     1.1,
		Description: "Create table urls",
		Script: `
CREATE TABLE urls (
	id       UUID,
	short_url        TEXT UNIQUE,
	long_url         TEXT ,
	date_created  TIMESTAMP,
	date_updated  TIMESTAMP,

	PRIMARY KEY (id)
);`,
	},
}

func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

const seeds = `
INSERT INTO urls (id, long_url, short_url,date_created, date_updated) VALUES
	('5cf37266-3473-4006-984f-9325122678b7', 'http://godoc.org/github.com/jmoiron/sqlx', 'iiWd', '2021-02-07 00:00:00', '2021-02-07 00:00:00'),
	('a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'https://www.reddit.com/', 'abcdef', '2021-02-07 00:00:00', '2021-02-07 00:00:00')
	ON CONFLICT DO NOTHING;
`

func DeleteDB(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(delete); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

const delete = `DELETE FROM urls;`
