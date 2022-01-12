package cohorts

import (
	"fmt"
	"reflect"
)

type Batch struct {
	Batch []Batches `json:"batch"`
}

type Batches struct {
	Type    string `json:"type"`
	Traits  Traits `json:"traits"`
	UserId  string `json:"userId"`
	Context struct {
		Integration struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"integration"`
	} `json:"context"`
}

type Traits map[string]interface{}

func (b Batches) GetCohortName() string {
	traitKey := reflect.ValueOf(b.Traits)

	var result []string

	for _, key := range traitKey.MapKeys() {
		result = append(result, key.String())
	}

	return result[0]
}

func A(b Batch) {
	fmt.Println(b)
}
