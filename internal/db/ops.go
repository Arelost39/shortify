package db

import (
	//"database/sql"
)

func (db *Database) GetLastID() {
	
}

func (db *Database) Encode(encoded_link string, decoded_link string) error {

	const q = `
		INSERT INTO shortify.encoded_links (decoded_link, encoded_link) VALUES ($1, $2)
	`

	_, err := db.pool.Exec(db.ctx, q, encoded_link, decoded_link)
	if err != nil {
		return err
	}
	
	return nil
}

func (db *Database) Decode(encoded_link string) (string, error) {
	
	const q = `
		SELECT decoded_link FROM
	`
	var result string
	return result, nil
}