package db

import "https://github.com/jmoiron/sqlx"

type Database struc {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode-%s"
		
	)	
}