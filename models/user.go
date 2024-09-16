package models

var HealthStatus = ""

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Job       string `json:"job"`
	Location  string `json:"location"`
	Email     string `json:"email"`
}
