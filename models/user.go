package models

var HealthStatus = "Server Runing"

type User struct {
	ID               int64  `json:"id"`
	FirstName        string `json:"ad"`
	LastName         string `json:"soyad"`
	Email            string `json:"ePosta"`
	GuncellemeTarihi string `json:"guncellemeTarihi"`
	TCKNO            string `json:"tcKimlikNo"`
}

type UserSource struct {
	Ad               string `json:"ad"`
	EPosta           string `json:"ePosta"`
	GuncellemeTarihi string `json:"guncellemeTarihi"`
	ID               int    `json:"id"`
	Soyad            string `json:"soyad"`
	TcKimlikNo       string `json:"tcKimlikNo"`
}

type UsersResponse struct {
	Users []User `json:"PersonelOzetDto"`
}

type UsersResponseSource struct {
	BagliKurulusPersonelListeSorgulaResult struct {
		Aciklama          string `json:"aciklama"`
		Data              []User `json:"data"`
		IslemBasarilimi   bool   `json:"islemBasarilimi"`
		ToplamKayitSayisi int    `json:"toplamKayitSayisi"`
	} `json:"BagliKurulusPersonelListeSorgulaResult"`
}

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
