package services

import (
	"bytes"
	"copy_users_for_moodle/models"
	"encoding/json"
	"net/http"
)

func CheckUserInLMS(apiURL string, email string) (bool, error) {
	resp, err := http.Get(apiURL + "?email=" + email)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}

func UpdateUserInLMS(apiURL string, user models.User) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, apiURL+"/"+user.Email, bytes.NewBuffer(userJson))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}
	return nil
}

func CreateUserInLMS(apiURL string, user models.User) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}
