package universities

// default data structure for api response
type UniversityInfo struct {
	AlphaTwoCode  string   `json:"alpha_two_code"`
	StateProvince string   `json:"state-province"`
	Domains       []string `json:"domains"`
	Country       string   `json:"country"`
	WebPages      []string `json:"web_pages"`
	Name          string   `json:"name"`
}
