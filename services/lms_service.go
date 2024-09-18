package services

import (
	"copy_users_for_moodle/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
func CheckUserInLMS(apiURL string, username string) (bool, error) {
	url := apiURL + "core_user_get_users&criteria[0][key]=username&moodlewsrestformat=json&criteria[0][value]=" + username
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}*/

func CheckUserInLMS(apiURL, lmsToken string, username string) (models.MoodleUser, error) {
	url := apiURL + "wstoken=" + lmsToken + "&wsfunction=core_user_get_users&criteria[0][key]=username&moodlewsrestformat=json&criteria[0][value]=" + username
	resp, err := http.Get(url)
	if err != nil {
		return models.MoodleUser{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.MoodleUser{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	// Read raw body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.MoodleUser{}, err
	}

	// Print raw JSON response for debugging
	//fmt.Println("Raw JSON response:", string(body))

	var apiResponse models.MoodleApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return models.MoodleUser{}, err
	}

	// Check if user was found
	if len(apiResponse.Users) > 0 {
		return apiResponse.Users[0], nil
	}

	return models.MoodleUser{}, nil
}

/*
func UpdateUserInLMS(apiURL string, user models.MoodleUser) error {
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
}*/

func UpdateUserInLMS(apiURL, lmsToken string, moodleUser models.MoodleUser, sourceUser models.User) error {
	//url := apiURL + "core_user_get_users&criteria[0][key]=username&moodlewsrestformat=json&criteria[0][value]=" + username

	updateBaseUrl := apiURL + "wstoken=" + lmsToken + "&wsfunction=core_user_update_users"
	updateUrl := createURLWithUserUpdate(moodleUser, sourceUser, updateBaseUrl)

	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, updateUrl, nil)

	if err != nil {
		//fmt.Println(err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	//fmt.Println("-")
	//fmt.Println("-")
	//fmt.Println(updateUrl)
	//fmt.Println("-")
	//fmt.Println("-")
	//fmt.Println(string(body))
	return nil
}

func CreateUserInLMS(apiURL, lmsToken string, sourceUser models.User) error {

	updateUrl := createURLWithUserCreate(sourceUser, apiURL+"wstoken="+lmsToken+"&wsfunction=core_user_create_users")

	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, updateUrl, nil)

	if err != nil {
		//fmt.Println(err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}

func createURLWithUserUpdate(moodleUser models.MoodleUser, sourceUser models.User, baseUrl string) string {
	// Suspended boolean değerini int'e dönüştürüyoruz (1: true, 0: false)
	suspended := 0
	if moodleUser.Suspended {
		suspended = 1
	}

	// URL'yi doğrudan fmt.Sprintf ile birleştiriyoruz
	finalURL := fmt.Sprintf("%s&users[0][email]=%s&users[0][firstname]=%s&users[0][id]=%d&users[0][lastname]=%s&users[0][suspended]=%d&users[0][username]=%s&moodlewsrestformat=json",
		baseUrl,
		url.QueryEscape(sourceUser.Email),     // Özel karakterleri escape et
		url.QueryEscape(sourceUser.FirstName), // İsimdeki özel karakterler için escape
		moodleUser.ID,
		url.QueryEscape(sourceUser.LastName), // Soyisimdeki özel karakterler için escape
		suspended,
		sourceUser.TCKNO)

	return finalURL
}

func createURLWithUserCreate(user models.User, baseUrl string) string {
	// URL'yi doğrudan fmt.Sprintf ile birleştiriyoruz
	finalURL := fmt.Sprintf("%s&users[0][username]=%s&users[0][password]=%s&users[0][firstname]=%s&users[0][lastname]=%s&users[0][email]=%s&moodlewsrestformat=json",
		baseUrl,
		url.QueryEscape(user.TCKNO), // Özel karakterleri escape et
		url.QueryEscape(user.TCKNO+"_"+firstN(user.FirstName)), // Şifreyi URL'ye uygun hale getir
		url.QueryEscape(user.FirstName),                        // İsimdeki özel karakterler için escape
		url.QueryEscape(user.LastName),                         // Soyisimdeki özel karakterler için escape
		url.QueryEscape(user.Email))                            // Email için escape

	return finalURL
}

func firstN(s string) string {
	n := 3
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
