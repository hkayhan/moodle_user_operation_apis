package controllers

import (
	"copy_users_for_moodle/models"
	"copy_users_for_moodle/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func CopyUsers(c *gin.Context) {

	usersAPI := os.Getenv("SOURCE_URL")
	sourceUserName := os.Getenv("SOURCE_USERNAME")
	sourcePassword := os.Getenv("SOURCE_PASSWORD")
	lmsAPI := os.Getenv("LMS_URL")
	lmsToken := os.Getenv("LMS_TOKEN")
	pageSize := 1000
	page := 1

	for {

		users, err := services.GetUsersPaginated(usersAPI, sourceUserName, sourcePassword, page, pageSize)
		if err != nil {
			fmt.Println("Kullanıcı verileri alınırken hata oluştu: %v", err)
		}

		if len(users) == 0 {
			fmt.Println("Tüm kullanıcılar işlendi.")
			break
		}

		for _, user := range users {

			if user.FirstName == "" {
				continue
			}
			lmsUser, err := services.CheckUserInLMS(lmsAPI, lmsToken, user.TCKNO)
			if err != nil {
				fmt.Println("LMS sorgusu sırasında hata oluştu: %v", err)
				continue
			}

			if lmsUser.ID > 0 {
				err := services.UpdateUserInLMS(lmsAPI, lmsToken, lmsUser, user)
				if err != nil {
					fmt.Println("Kullanıcı güncellenirken hata oluştu: %v", err)
				} else {
					fmt.Println("Kullanıcı güncellendi:", user.TCKNO)
				}

			} else {
				err := services.CreateUserInLMS(lmsAPI, lmsToken, user)
				if err != nil {
					fmt.Println("Yeni kullanıcı oluşturulurken hata oluştu: %v", err)
				} else {
					fmt.Println("Yeni kullanıcı oluşturuldu: ", user.TCKNO)
				}
			}

			fmt.Println(lmsUser)
			/*
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
				}*/
		}

		// Bir sonraki sayfaya geçiyoruz
		page++
	}

}

func TestGo(c *gin.Context) {

	c.String(200, models.HealthStatus)
}
