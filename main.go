package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("user.db.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	db := Userdb{}
	dec.Decode(&db)

	fmt.Println(db)
	// createJSONFile(db)
}

// User is a structure for the json files
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Userdb is a database of users
type Userdb struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}

func createJSONFile(db Userdb) {
	users := []User{
		{Username: "TonyMoon", Password: "TonysPass", Email: "tonymoon@gmail.com"},
		{Username: "RosaCruz", Password: "RosasPass", Email: "rosaacruz@gmail.com"},
		{Username: "FlorLuna", Password: "FlorsPass", Email: "florluna@gmail.com"},
	}
	db = Userdb{Users: users, Type: "Simple"}

	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(db)
	f, err := os.Create("user.db.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
