package models
import "fmt"
import (
	"errors"
	"fmt"
)

type User struct {
	ID      int
	Name    string // must be Name not name :D, won't work
	Age     int
	Address Address
}


	// I removed the in-memory database of users, you will use the disk
	// create a file for each added user
	// when a user is requested by id you should look for it's file to return it
	nextID = NUM_OF_FILES // NUM_OF_FILES: should reflect the number of files in users_saved

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
	return u, nil
}

func GetUserByID(id int) (User, error) {
	/*
		TODO
		1) look for the file named "u.ID".txt in users_saved directory
		2) unmarshal the read json into a user and return that user with nil error
	*/


	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

