package controller

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"opc-site/app/model"
)

type UserRepository struct {
	SQLHandler SQLHandler
}

const usersTableName = "users"

func (ur *UserRepository) Add(user *model.User) bool {
	query := "SELECT * FROM " + usersTableName + " WHERE name=?"
	userInDb := &model.User{}
	err := ur.SQLHandler.Db.QueryRow(
		query,
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
		ur.ConvertPassword(user.Password),
	)
	if err != nil {
		log.Panic(err.Error())
	}

	return true
}

func (ur *UserRepository) Authorize(user *model.User) bool {
	user.Password = ur.ConvertPassword(user.Password)

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

func (ur *UserRepository) GetUserById(userId int) *model.User {
	user := &model.User{}
	query := "SELECT * FROM " + usersTableName + " WHERE id=?"
	err := ur.SQLHandler.Db.QueryRow(
		query,
		userId,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Password,
	)
	if err != nil {
		log.Println("No user with ID")
		return user
	}

	return user
}

func (ur *UserRepository) ConvertPassword(password string) string {
	passwordHash := md5.Sum([]byte(password))
	return hex.EncodeToString(passwordHash[:])
}
