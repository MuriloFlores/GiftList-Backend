package model

type LoginResponse struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Token       string `json:"token"`
}
