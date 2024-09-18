package services

import (
	"copy_users_for_moodle/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
func GetUsersPaginated(apiURL string, page int, pageSize int) ([]models.User, error) {

		jsonResponse := `
	    [
	    {
	      "id": 1,
	      "firstName": "Test Update",
	      "lastName": "Test Update LastName",
	      "email": "test_update@example.com",
	      "TCKNO": "28984632406"
	    },
	    {
	      "id": 2,
	      "firstName": "Jane",
	      "lastName": "Smith",
	      "email": "jane.smith@example.com",
	      "TCKNO": "21111111112"
	    },
	    {
	      "id": 3,
	      "firstName": "Michael",
	      "lastName": "Johnson",
	      "email": "michael.johnson@example.com",
	      "TCKNO": "21111111113"
	    },
	    {
	      "id": 4,
	      "firstName": "Emily",
	      "lastName": "Williams",
	      "email": "emily.williams@example.com",
	      "TCKNO": "21111111114"
	    },
	    {
	      "id": 5,
	      "firstName": "David",
	      "lastName": "Brown",
	      "email": "david.brown@example.com",
	      "TCKNO": "21111111115"
	    },
	    {
	      "id": 6,
	      "firstName": "Sophia",
	      "lastName": "Miller",
	      "email": "sophia.miller@example.com",
	      "TCKNO": "21111111116"
	    },
	    {
	      "id": 7,
	      "firstName": "James",
	      "lastName": "Davis",
	      "email": "james.davis@example.com",
	      "TCKNO": "21111111117"
	    },
	    {
	      "id": 8,
	      "firstName": "Olivia",
	      "lastName": "Garcia",
	      "email": "olivia.garcia@example.com",
	      "TCKNO": "21111111118"
	    },
	    {
	      "id": 9,
	      "firstName": "Christopher",
	      "lastName": "Rodriguez",
	      "email": "christopher.rodriguez@example.com",
	      "TCKNO": "21111111119"
	    },
	    {
	      "id": 10,
	      "firstName": "Isabella",
	      "lastName": "Martinez",
	      "email": "isabella.martinez@example.com",
	      "TCKNO": "21111111120"
	    },
	    {
	      "id": 11,
	      "firstName": "Daniel",
	      "lastName": "Hernandez",
	      "email": "daniel.hernandez@example.com",
	      "TCKNO": "21111111121"
	    },
	    {
	      "id": 12,
	      "firstName": "Grace",
	      "lastName": "Lopez",
	      "email": "grace.lopez@example.com",
	      "TCKNO": "21111111122"
	    },
	    {
	      "id": 13,
	      "firstName": "Matthew",
	      "lastName": "Gonzalez",
	      "email": "matthew.gonzalez@example.com",
	      "TCKNO": "21111111123"
	    },
	    {
	      "id": 14,
	      "firstName": "Sofia",
	      "lastName": "Wilson",
	      "email": "sofia.wilson@example.com",
	      "TCKNO": "21111111124"
	    },
	    {
	      "id": 15,
	      "firstName": "Joshua",
	      "lastName": "Anderson",
	      "email": "joshua.anderson@example.com",
	      "TCKNO": "21111111125"
	    },
	    {
	      "id": 16,
	      "firstName": "Emma",
	      "lastName": "Thomas",
	      "email": "emma.thomas@example.com",
	      "TCKNO": "21111111126"
	    },
	    {
	      "id": 17,
	      "firstName": "Alexander",
	      "lastName": "Taylor",
	      "email": "alexander.taylor@example.com",
	      "TCKNO": "21111111127"
	    },
	    {
	      "id": 18,
	      "firstName": "Mia",
	      "lastName": "Moore",
	      "email": "mia.moore@example.com",
	      "TCKNO": "21111111128"
	    },
	    {
	      "id": 19,
	      "firstName": "Samuel",
	      "lastName": "Jackson",
	      "email": "samuel.jackson@example.com",
	      "TCKNO": "21111111129"
	    },
	    {
	      "id": 20,
	      "firstName": "Charlotte",
	      "lastName": "White",
	      "email": "charlotte.white@example.com",
	      "TCKNO": "21111111130"
	    },
	    {
	      "id": 21,
	      "firstName": "Joseph",
	      "lastName": "Harris",
	      "email": "joseph.harris@example.com",
	      "TCKNO": "21111111131"
	    },
	    {
	      "id": 22,
	      "firstName": "Amelia",
	      "lastName": "Martin",
	      "email": "amelia.martin@example.com",
	      "TCKNO": "21111111132"
	    },
	    {
	      "id": 23,
	      "firstName": "Anthony",
	      "lastName": "Thompson",
	      "email": "anthony.thompson@example.com",
	      "TCKNO": "21111111133"
	    },
	    {
	      "id": 24,
	      "firstName": "Harper",
	      "lastName": "Martinez",
	      "email": "harper.martinez@example.com",
	      "TCKNO": "21111111134"
	    },
	    {
	      "id": 25,
	      "firstName": "Andrew",
	      "lastName": "Walker",
	      "email": "andrew.walker@example.com",
	      "TCKNO": "21111111135"
	    },
	    {
	      "id": 26,
	      "firstName": "Evelyn",
	      "lastName": "Young",
	      "email": "evelyn.young@example.com",
	      "TCKNO": "21111111136"
	    },
	    {
	      "id": 27,
	      "firstName": "Ryan",
	      "lastName": "King",
	      "email": "ryan.king@example.com",
	      "TCKNO": "21111111137"
	    },
	    {
	      "id": 28,
	      "firstName": "Ella",
	      "lastName": "Scott",
	      "email": "ella.scott@example.com",
	      "TCKNO": "21111111138"
	    },
	    {
	      "id": 29,
	      "firstName": "Benjamin",
	      "lastName": "Torres",
	      "email": "benjamin.torres@example.com",
	      "TCKNO": "21111111139"
	    },
	    {
	      "id": 30,
	      "firstName": "Avery",
	      "lastName": "Nguyen",
	      "email": "avery.nguyen@example.com",
	      "TCKNO": "21111111140"
	    },
	    {
	      "id": 31,
	      "firstName": "Logan",
	      "lastName": "Hill",
	      "email": "logan.hill@example.com",
	      "TCKNO": "21111111141"
	    },
	    {
	      "id": 32,
	      "firstName": "Scarlett",
	      "lastName": "Flores",
	      "email": "scarlett.flores@example.com",
	      "TCKNO": "21111111142"
	    },
	    {
	      "id": 33,
	      "firstName": "Ethan",
	      "lastName": "Green",
	      "email": "ethan.green@example.com",
	      "TCKNO": "21111111143"
	    },
	    {
	      "id": 34,
	      "firstName": "Zoey",
	      "lastName": "Adams",
	      "email": "zoey.adams@example.com",
	      "TCKNO": "21111111144"
	    },
	    {
	      "id": 35,
	      "firstName": "Luke",
	      "lastName": "Nelson",
	      "email": "luke.nelson@example.com",
	      "TCKNO": "21111111145"
	    },
	    {
	      "id": 36,
	      "firstName": "Hannah",
	      "lastName": "Baker",
	      "email": "hannah.baker@example.com",
	      "TCKNO": "21111111146"
	    },
	    {
	      "id": 37,
	      "firstName": "Jack",
	      "lastName": "Rivera",
	      "email": "jack.rivera@example.com",
	      "TCKNO": "21111111147"
	    }
	  ]`

		var users models.UsersResponse

		// JSON'u struct'a çevirme
		err := json.Unmarshal([]byte(jsonResponse), &users)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return []models.User{}, err
		}

		// Veriyi yazdırma
		if page == 1 {
			return users, err

		}
		return []models.User{}, nil
	}
*/
func GetUsersPaginated(apiUrl, sourceUserName, sourcePassword string, page, pageSize int) ([]models.User, error) {

	url := apiUrl
	method := "POST"

	payload := strings.NewReader(`{` + "\n" + `
	"sayfaNo":` + strconv.Itoa(page) + ",\n" + `"sayfaVeriSayisi":` + strconv.Itoa(pageSize) + "\n" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("username", sourceUserName)
	req.Header.Add("userpassword", sourcePassword)
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Cookie", "ASP.NET_SessionId=hzyjqxsgxbxqkgxbe4rnpv2o")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(string(body))

	//var apiResponse models.UsersResponse

	var icisleriResponse models.UsersResponseSource
	if err := json.Unmarshal(body, &icisleriResponse); err != nil {
		return nil, err
	}

	return icisleriResponse.BagliKurulusPersonelListeSorgulaResult.Data, nil
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
