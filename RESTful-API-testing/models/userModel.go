package models

type UserModel interface {
	Get() []User
	Insert(User) error
}
