package cohorts

import "reflect"

type Batch struct {
	Batch []Batches `json:"batch"`
}

type Batches struct {
	Type    string                 `json:"type"`
	Traits  map[string]interface{} `json:"traits"`
	UserId  string                 `json:"userId"`
	Context struct {
		Integration struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"integration"`
	} `json:"context"`
}

func GetCohortName(trait interface{}) string {
	traitKey := reflect.ValueOf(trait)

	var result []string

	for _, key := range traitKey.MapKeys() {
		result = append(result, key.String())
	}

	return result[0]
}
