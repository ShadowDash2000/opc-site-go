package controller

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"log"
	"opc-site/app/model"
)

type UserRepository struct {
	SQLHandler SQLHandler
}

func (ur *UserRepository) Add(user *model.User) bool {
	query := "SELECT * FROM ? WHERE name=?"
	userInDb := &model.User{}
	err := ur.SQLHandler.Db.QueryRow(
		query,
		usersTableName,
		user.Name,
	).Scan(
		&userInDb.Id,
		&userInDb.Name,
		&userInDb.Password,
	)
	if err == nil {
		return false
	}

	query = "INSERT INTO " + usersTableName + " (name, password) VALUES (?, ?)"
	_, err = ur.SQLHandler.Db.Exec(
		query,
		user.Name,
		md5.Sum([]byte(user.Password)),
	)
	if err != nil {
		log.Panic(err.Error())
	}

	return true
}

func (ur *UserRepository) Authorize(user *model.User) bool {
	query := "SELECT * FROM " + usersTableName + " WHERE name=? AND password=?"
	err := ur.SQLHandler.Db.QueryRow(
		query,
		user.Name,
		user.Password,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Password,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	} else if err != nil {
		log.Panic(err)
	}

	return true
}

func (ur *UserRepository) Logout() bool {

	return true
}
