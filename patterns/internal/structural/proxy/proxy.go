package proxy

import "fmt"

// Database representation.
// It contains info about users.
type UserListProxy struct {
	DB                     UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int64) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	} else {
		fmt.Println("Returning user from database")
		u.addUserToStack(user)
		u.DidLastSearchUsedCache = false
		return user, nil
	}
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (ul *UserList) addUser(u User) {
	*ul = append(*ul, u)
}
