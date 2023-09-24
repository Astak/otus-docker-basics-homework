package data

import (
	"fmt"
	"sync"
)

type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UserStore struct {
	sync.Mutex

	users  map[int64]User
	nextId int64
}

func NewUserStore() *UserStore {
	us := &UserStore{}
	us.users = make(map[int64]User)
	us.nextId = 0
	return us
}

func (store *UserStore) CreateUser(userName string, firstName string, lastName string, email string, phone string) User {
	store.Lock()
	defer store.Unlock()
	user := User{
		ID:        store.nextId,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	store.users[store.nextId] = user
	store.nextId++
	return user
}

func (us *UserStore) GetUser(id int64) (User, error) {
	us.Lock()
	defer us.Unlock()
	u, ok := us.users[id]
	if ok {
		return u, nil
	} else {
		return User{}, fmt.Errorf("user with id=%d not found", id)
	}
}

func (us *UserStore) DeleteUser(id int64) error {
	us.Lock()
	defer us.Unlock()

	if _, ok := us.users[id]; !ok {
		return fmt.Errorf("user with id=%d not found", id)
	}
	delete(us.users, id)
	return nil
}

func (store *UserStore) UpdateUser(id int64, userName string, firstName string, lastName string, email string, phone string) (User, error) {
	store.Lock()
	defer store.Unlock()
	user, ok := store.users[id]
	if ok {
		user.UserName = userName
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.Phone = phone
		store.users[id] = user
		return user, nil
	} else {
		return User{}, fmt.Errorf("user with id=%d not found", id)
	}
}
