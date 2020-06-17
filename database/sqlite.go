package database

import "database/sql"

type Database struct {
	Dir string
}

func (d Database)Connect() (*sql.DB,error){

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil,err
	}
	defer db.Close()

	return db, nil
}
