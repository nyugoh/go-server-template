package models

type User struct {
	UserId uint64 `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
