package models

type LoginAuth struct {
	Email        string             `json:"email"`
	Password string             `json:"hash_password"`
}
