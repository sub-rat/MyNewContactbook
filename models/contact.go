package models

type Phone struct {
	PhoneType   string `json:"type"`
	PhoneNumber string `json:"phone_number"`
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type Contact struct {
	ID        uint    `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     []Phone `json:"phone"`
	Address   Address `json:"address"`
}
