package database

import (
	"backend/internal/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func NewConnection(s *config.DatabaseConfig) error {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.DbName)

	db, err := sqlx.Connect("postgres", dataSource)
	s.Db = db

	if err != nil {
		log.Println("Connect database failed")
		return err
	}

	log.Println("Connect database succesfully!")
	return nil
}
