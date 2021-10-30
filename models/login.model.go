package models

import (
	"database/sql"
	"echo-rest/db"
	"echo-rest/helper"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Username atau Password salah")
		return false, err
	}

	return true, nil
}
