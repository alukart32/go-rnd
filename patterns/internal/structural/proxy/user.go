package proxy

import "fmt"

type UserFinder interface {
	FindUser(id int64) (User, error)
}

type User struct {
	ID int64
}

type UserList []User

func (u *UserList) FindUser(id int64) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{ID: id}, fmt.Errorf("user %d could not be found", id)
}
