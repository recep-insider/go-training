package upsert

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Identifiers struct {
		Email string `json:"email"`
	} `json:"identifiers"`
	Attributes struct {
		Custom struct {
			Acn [1]string `json:"acn"`
		} `json:"custom"`
	} `json:"attributes"`
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

type CombinedResult struct {
	Insert map[string]interface{} `json:"insert"`
	Delete map[string]interface{} `json:"delete"`
}
