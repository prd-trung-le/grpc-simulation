package repository

import (
	"auth/models"
	"database/sql"
	"log"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	CreateUser(username string, password string) (bool, error)
	GetUserByUsername(username string) (models.User, error)
}

func InitUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) CreateUser(username string, password string) (bool, error) {
	var err error
	var result bool

	tx, errTx := userRepository.db.Begin()
	if errTx != nil {
		log.Println("Error create user: ", errTx)
	} else {
		err = insertUser(tx, username, password)
		if err != nil {
			log.Println("Error create user: ", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error create user: ", err)
	}

	return result, err
}

func insertUser(tx *sql.Tx, username string, password string) error {
	raw_query := fmt.Sprintf("INSERT INTO user (username, password) VALUES('%s', '%s');", username, password)
	res, err := tx.Exec(raw_query)
	fmt.Println(res)
	return err
}

func (userRepository *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	var id int64

	err := userRepository.db.QueryRow(`
		SELECT id, username, password FROM user WHERE username=%s;
	`, username).Scan(&id, &(user.Username), &(user.Password))

	user.ID = id

	if err != nil {
		log.Println("Error to get user by username", err)
	}

	return user, err
}
