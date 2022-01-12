package upsert

import (
	"bytes"
	"encoding/json"
	"go-training/internal/cohorts"
	"io/ioutil"
	"net/http"
)

const (
	apiUrl string = "https://unification.useinsider.com/api/user/v1/"
)

type InsertUsers struct {
	Users []InsertUser `json:"users"`
}

type InsertUser struct {
	Identifiers struct {
		Email string `json:"email"`
	} `json:"identifiers"`
	Attributes struct {
		Custom struct {
			Acn [1]string `json:"acn"`
		} `json:"custom"`
	} `json:"attributes"`
}

func GetInsertUsers(batch cohorts.Batch) string {
	user := InsertUser{}
	users := InsertUsers{}

	for _, item := range batch.Batch {
		user.Identifiers.Email = item.UserId

		if item.Traits[item.GetCohortName()].(bool) {
			user.Attributes.Custom.Acn[0] = item.GetCohortName()
			users.Users = append(users.Users, user)
		}
	}

	userJson, _ := json.Marshal(users)

	return string(userJson)
}

type DeleteUsers struct {
	Users []DeleteUser `json:"users"`
}

type DeleteUser struct {
	Identifiers struct {
		Email string `json:"email"`
	} `json:"identifiers"`
	Custom struct {
		Partial struct {
			Acn [1]string `json:"acn"`
		} `json:"partial"`
	} `json:"custom"`
}

func GetDeleteUsers(batch cohorts.Batch) string {
	user := DeleteUser{}
	users := DeleteUsers{}

	for _, item := range batch.Batch {
		user.Identifiers.Email = item.UserId

		if !item.Traits[item.GetCohortName()].(bool) {
			user.Custom.Partial.Acn[0] = item.GetCohortName()
			users.Users = append(users.Users, user)
		}
	}

	userJson, _ := json.Marshal(users)

	return string(userJson)
}

type CombinedResult struct {
	Insert map[string]interface{} `json:"insert"`
	Delete map[string]interface{} `json:"delete"`
}

func SendUpsertRequest(url, payload string) map[string]interface{} {
	request, _ := http.NewRequest("POST", apiUrl+url, bytes.NewBuffer([]byte(payload)))

	request.Header.Set("X-PARTNER-NAME", "yigittest")
	request.Header.Set("X-REQUEST-TOKEN", "b520d6eb6a035985b3c354d30c9b855da069ae68a2f0c18ffa5c6b6a991a76")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, responseError := client.Do(request)

	if responseError != nil {
		panic(responseError)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var jsonError map[string]interface{}

	json.Unmarshal([]byte(string(body)), &jsonError)

	return jsonError
}
