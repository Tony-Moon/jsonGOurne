package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	users := []User{
		{
			Username: "TonyMoon",
			Password: "TonysPass",
			Email:    "tonymoon@gmail.com",
		},
		{
			Username: "RosaCruz",
			Password: "RosasPass",
			Email:    "rosaacruz@gmail.com",
		},
		{
			Username: "FlorLuna",
			Password: "FlorsPass",
			Email:    "florluna@gmail.com",
		},
	}

	db := Userdb{
		Users: users,
		Type:  "simple",
	}

	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(db)

	f, err := os.Create("user.db.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	io.Copy(os.Stdout, buf)
}

// User is a structure for the json files
type User struct {
	Username string
	Password string
	Email    string
}

// Userdb is a database of users
type Userdb struct {
	Users []User
	Type  string
}
