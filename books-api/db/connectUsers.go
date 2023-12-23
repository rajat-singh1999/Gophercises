package db

import (
	"encoding/json"
	"errors"
	Config "example/books-api/config"
	Structs "example/books-api/types"
	"io/ioutil"
	"os"
)

type User = Structs.User

var userPath = Config.UserPath()

func GetUser(id string) (User, error) {
	var temp User
	if users, err := LoadUsersFromFile(); err != nil {
		return temp, err
	} else {
		for i, b := range users {
			if b.ID == id {
				return users[i], nil
			}
		}
		return temp, errors.New("user not found")
	}
}

func GetUserByUsername(username string) (*User, error) {
	if users, err := LoadUsersFromFile(); err != nil {
		return nil, err
	} else {
		for i, b := range users {
			if b.Username == username {
				return &users[i], nil
			}
		}
		return nil, errors.New("user not found")
	}
}

func GetUserByEmail(email string) (*User, error) {
	if users, err := LoadUsersFromFile(); err != nil {
		return nil, err
	} else {
		for i, b := range users {
			if b.Email == email {
				return &users[i], nil
			}
		}
		return nil, errors.New("user not found")
	}
}

func LoadUsersFromFile() ([]User, error) {
	jsonfile, err := os.Open(userPath)
	if err != nil {
		return nil, errors.New("trouble reading hson file")
	}
	defer jsonfile.Close()

	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		return nil, err
	}
	var users []User
	json.Unmarshal(byteValue, &users)
	return users, nil
}

func SaveUsersToFile(users []User) error {
	file, err := os.Create(userPath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func SaveUserToUsers(user User) ([]User, error) {
	err := DeleteUser(user.ID)
	if err != nil {
		return nil, err
	}

	users, err := LoadUsersFromFile()
	if err != nil {
		return nil, err
	}

	users = append(users, user)

	err = SaveUsersToFile(users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UserReturnsBook(uid string, bid string) (User, error, error) {
	user, err := GetUser(uid)
	if err != nil {
		return user, nil, err
	}

	issued := user.Issued
	var indexToDelete int = -1
	for i, b := range issued {
		if bid == b {
			indexToDelete = i
			break
		}
	}
	if indexToDelete == -1 {
		return user, nil, errors.New("the user did not issue this book! Bad request")
	}

	user.Issued = append(issued[:indexToDelete], issued[indexToDelete+1:]...)
	_, err = SaveUserToUsers(user)
	if err != nil {
		return user, err, nil
	}

	return user, nil, nil
}

func DeleteUser(id string) error {
	if users, err := LoadUsersFromFile(); err != nil {
		return err
	} else {
		for i, b := range users {
			if b.ID == id {
				users = append(users[:i], users[i+1:]...)
				if err = SaveUsersToFile(users); err != nil {
					return errors.New("error while saving to file")
				} else {
					return nil
				}
			}
		}
		return errors.New("no book matching the passed ID found")
	}
}

func CheckIfBookIssuedByUser(uid string, bid string) error {
	user, err := GetUser(uid)
	if err != nil {
		return errors.New("the user does not exist")
	}

	issued := user.Issued
	for _, v := range issued {
		if v == bid {
			return errors.New("the user has already issued this book")
		}
	}
	return nil
}
