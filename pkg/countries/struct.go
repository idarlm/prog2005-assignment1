package countries

// partial implementation of countries response
type BasicInfo struct {
	Name      Names             `json:"name"`
	Capital   []string          `json:"capital"`
	Languages map[string]string `json:"languages"`
	Borders   []string          `json:"borders"`
	Maps      MapsLinks         `json:"maps"`
	Cca2      string            `json:"cca2"`
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
