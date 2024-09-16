package controllers

import (
	"copy_users_for_moodle/models"
	"copy_users_for_moodle/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func CopyUsers(c *gin.Context) {

	usersAPI := os.Getenv("SOURCE_URL") + "/users"
	lmsAPI := os.Getenv("MOODLE_URL") + "/users"
	pageSize := 1000 // her seferde 1000 kullanıcı çekeceğiz
	page := 1        // ilk sayfadan başlıyoruz

	for {

		users, err := services.GetUsersPaginated(usersAPI, page, pageSize)
		if err != nil {
			fmt.Println("Kullanıcı verileri alınırken hata oluştu: %v", err)
		}

		if len(users) == 0 {
			// Hiç kullanıcı yoksa veya sayfa boşsa, döngüden çıkıyoruz
			fmt.Println("Tüm kullanıcılar işlendi.")
			break
		}

		for _, user := range users {
			exists, err := services.CheckUserInLMS(lmsAPI, user.Email)
			if err != nil {
				fmt.Println("LMS sorgusu sırasında hata oluştu: %v", err)
				continue
			}

			if exists {
				err := services.UpdateUserInLMS(lmsAPI, user)
				if err != nil {
					fmt.Println("Kullanıcı güncellenirken hata oluştu: %v", err)
				} else {
					fmt.Println("Kullanıcı güncellendi: %s\n", user.Email)
				}
			} else {
				err := services.CreateUserInLMS(lmsAPI, user)
				if err != nil {
					fmt.Println("Yeni kullanıcı oluşturulurken hata oluştu: %v", err)
				} else {
					fmt.Println("Yeni kullanıcı oluşturuldu: ", user.Email)
				}
			}
		}

		// Bir sonraki sayfaya geçiyoruz
		page++
	}

}

func TestGo(c *gin.Context) {

	c.String(200, models.HealthStatus)
}
