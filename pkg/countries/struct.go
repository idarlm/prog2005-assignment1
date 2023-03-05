package countries

// partial implementation of countries response
type BasicInfo struct {
	Name    Names     `json:"name"`
	Capital []string  `json:"capital"`
	Borders []string  `json:"borders"`
	Maps    MapsLinks `json:"maps"`
}

// data type containing language fields
type Language struct {
	Iso639_1   string `json:"iso639_1"`
	Iso639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

// data type containing name fields
type Names struct {
	Common      string           `json:"common"`
	Official    string           `json:"official"`
	NativeNames map[string]Names `json:"nativeName"`
}

// data type containing map fields
type MapsLinks struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}
