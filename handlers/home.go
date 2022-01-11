package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pergamon/internals/cohorts"
	"pergamon/internals/upsert"
	"reflect"

	"github.com/labstack/echo/v4"
)

const (
	apiUrl string = "https://unification.useinsider.com/api/user/v1/"
)

func Home(c echo.Context) error {
	batch := cohorts.Batch{}
	user := upsert.User{}
	users := upsert.Users{}
	deleteUser := upsert.DeleteUser{}
	deleteUsers := upsert.DeleteUsers{}

	err := c.Bind(&batch)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	json.NewDecoder(c.Request().Body).Decode(&batch)

	for _, item := range batch.Batch {
		user.Identifiers.Email = item.UserId
		deleteUser.Identifiers.Email = item.UserId

		if item.Traits[GetCohortName(item.Traits)].(bool) {
			user.Attributes.Custom.Acn[0] = GetCohortName(item.Traits)
			users.Users = append(users.Users, user)
		} else {
			deleteUser.Custom.Partial.Acn[0] = GetCohortName(item.Traits)
			deleteUsers.Users = append(deleteUsers.Users, deleteUser)
		}
	}

	upsertPayload, _ := json.Marshal(users)
	upsertResult := SendUpsertRequest("upsert", string(upsertPayload))

	deletePayload, _ := json.Marshal(deleteUsers)
	deleteResult := SendUpsertRequest("attribute/delete", string(deletePayload))

	return c.JSON(http.StatusOK, upsert.CombinedResult{
		Insert: upsertResult,
		Delete: deleteResult,
	})
}

func GetCohortName(trait interface{}) string {
	traitKey := reflect.ValueOf(trait)

	var result []string

	for _, key := range traitKey.MapKeys() {
		result = append(result, key.String())
	}

	return result[0]
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
