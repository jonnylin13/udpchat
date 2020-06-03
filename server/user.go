package server

import "net"

// User represents a user.
type User struct {
	name string
	addr net.Addr
}

// GetNames returns a []string of names from []User.
func GetNames(userList []User) []string {
	var list []string
	for _, user := range userList {
		list = append(list, user.name)
	}
	return list
}

// RemoveUser returns a []User with the removed name.
func RemoveUser(userList []User, name string) []User {
	for i, user := range userList {
		if user.name == name {
			return append(userList[:i], userList[i+1:]...)
		}
	}
	return userList
}
