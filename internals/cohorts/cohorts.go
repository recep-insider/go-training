package cohorts

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
