package models

var HealthStatus = ""

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Job       string `json:"job"`
	TCKNO     string `json:"TCKNO"`
}

type UsersResponse []User

type MoodleUser struct {
	ID                   int    `json:"id"`
	Username             string `json:"username"`
	Firstname            string `json:"firstname"`
	Lastname             string `json:"lastname"`
	Fullname             string `json:"fullname"`
	Email                string `json:"email"`
	Department           string `json:"department"`
	Firstaccess          int    `json:"firstaccess"`
	Lastaccess           int    `json:"lastaccess"`
	Auth                 string `json:"auth"`
	Suspended            bool   `json:"suspended"`
	Confirmed            bool   `json:"confirmed"`
	Lang                 string `json:"lang"`
	Theme                string `json:"theme"`
	Timezone             string `json:"timezone"`
	Mailformat           int    `json:"mailformat"`
	City                 string `json:"city"`
	Profileimageurl      string `json:"profileimageurl"`
	Profileimageurlsmall string `json:"profileimageurlsmall"`
}

type MoodleApiResponse struct {
	Users    []MoodleUser `json:"users"`
	Warnings []string     `json:"warnings"`
}
