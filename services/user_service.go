package services

import (
	"copy_users_for_moodle/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetUsersPaginated(apiURL string, page int, pageSize int) ([]models.User, error) {
	url := fmt.Sprintf("%s?page=%d&pageSize=%d", apiURL, page, pageSize)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersCountFromSource() (int, error) {

	sourceUrl := os.Getenv("SOURCE_URL")
	resp, err := http.Get(sourceUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return 0, err
	}
	return 0, nil
}
