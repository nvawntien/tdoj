package db

import (
	"backend/internal/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func NewConnection(s *config.DatabaseConfig) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.DbName)

	db, err := sqlx.Connect("postgres", dataSource)

	if err != nil {
		log.Println("Connect database failed")
		return nil, err
	}

	log.Println("Connect database succesfully!")
	return db, nil
}
