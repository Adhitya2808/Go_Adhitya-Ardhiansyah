package main

type User struct {
	ID       int
	Username string
	Password string
}

type UserService struct {
	users []User
}

func (s UserService) GetAllUsers() []User {
	return s.users
}

func (s UserService) GetUserByID(id int) User {
	for _, user := range s.users {
		if id == user.ID {
			return user
		}
	}

	return User{}
}