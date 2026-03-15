package db

import (
	//"database/sql"
)

func (db *Database) GetLastID() (string, error) {
	
	const q = `
		SELECT COALESCE(MAX(id), 0) FROM encoded_links;
	`
	rows, err := db.pool.Query(db.ctx, q)
	if err != nil {
		return "", err
	}
	var result string

	if err = rows.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}

func (db *Database) Encode(encoded_link string, decoded_link string) error {

	const q = `
		INSERT INTO shortify.encoded_links (decoded_link, encoded_link) VALUES ($1, $2);
	`

	_, err := db.pool.Exec(db.ctx, q, encoded_link, decoded_link)
	if err != nil {
		return err
	}
	
	return nil
}

func (db *Database) Decode(encoded_link string) (string, error) {

	const q = `
		SELECT * FROM shortify.encoded_links WHERE encoded_link = $1;
	`

	rows, err := db.pool.Query(db.ctx, q, encoded_link)
	if err != nil {
		return "", err
	}

	var result string

	if err = rows.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}