package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type User struct {
	ID      int
	Name    string // must be Name not name :D, won't work
	Age     int
	Address Address
}

func get_num_files(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print("Error")
	}
	return len(files)
}

var (
	NUM_OF_FILES int = get_num_files("users_saved")
	nextID           = NUM_OF_FILES + 1
)

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	/*
		TODO
		1) create a file named as the "u.ID".txt and save into users_saved directory
		2) marshal the user's json into it
	*/

	f, err := os.Create("users_saved/" + strconv.Itoa(u.ID) + ".txt")
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(u)
	if err != nil {
		f.Close()
		return User{}, err
	}
	f.WriteString(string(data))
	f.Close()

	return u, nil
}

func GetUserByID(id int) (User, error) {
	/*
		TODO
		1) look for the file named "u.ID".txt in users_saved directory
		2) unmarshal the read json into a user and return that user with nil error
	*/

	stream, err := ioutil.ReadFile("users_saved/" + strconv.Itoa(id) + ".txt")
	if err != nil {
		panic(err)
	}
	var u User
	json.Unmarshal(stream, &u)

	return u, fmt.Errorf("User with ID '%v' not found", id)
}
