package models

import (
	"errors"
	"fmt"

	"github.com/prezessikora/events/db"
	"github.com/prezessikora/events/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) VerifyCredetials() error {
	query := `SELECT id, password FROM users WHERE email=?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	var userId int64
	err := row.Scan(&userId, &retrievedPassword)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Veryfied credentails for userid: %v", userId)
	user.ID = userId
	passwordIsVaild := utils.CheckHashedPassword(user.Password, retrievedPassword)
	if !passwordIsVaild {
		fmt.Println("credentials invalid")
		return errors.New("credentials invalid")
	}

	return nil

}

func (user *User) Save() error {

	insertSql := `INSERT INTO users (email,password)
	VALUES (?,?)`
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	user.ID, err = result.LastInsertId()
	return err
}
