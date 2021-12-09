package domain

import (
	// postgres driver

	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/weeee9/hexagonal-architecture/config"

	"github.com/jmoiron/sqlx"
)

func NewCustomerRepotitoryPostgres() CustomerRepositoryPostgres {
	cfg, err := config.Environ()
	if err != nil {
		log.Fatalf("fail to init config: %v", err.Error())
	}

	db, err := sqlx.Open("postgres", getPostgresConnetionInfo(cfg))
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryPostgres{
		db: db,
	}
}

func getPostgresConnetionInfo(cfg config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
}

type CustomerRepositoryPostgres struct {
	db *sqlx.DB
}

func (repo CustomerRepositoryPostgres) FindAll() ([]Customer, error) {
	querySQL := `SELECT id, name, city FROM customer`

	log.Println(querySQL)
	rows, err := repo.db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cusomters := []Customer{}
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.City); err != nil {
			log.Printf("error scanning customer %s\n", err.Error())
			return nil, err
		}
		cusomters = append(cusomters, customer)
	}

	return cusomters, nil
}

func (repo CustomerRepositoryPostgres) GetByID(id int) (*Customer, error) {
	querySQL := `SELECT id, name, city FROM customer WHERE id = $1`

	log.Printf("%s [%v]\n", querySQL, id)
	row := repo.db.QueryRow(querySQL, id)

	customer := Customer{}
	if err := row.Scan(&customer.ID, &customer.Name, &customer.City); err != nil {
		log.Printf("error finding customer %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found with id: %v", id)
		}

		return nil, err
	}

	return &customer, nil
}
