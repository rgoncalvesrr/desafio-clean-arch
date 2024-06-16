package database

import "database/sql"

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
	create table if not exists orders (
		id varchar(15),
		price double default 0,
		tax double default 0,
		final_price double default 0
	);`)

	return err
}
