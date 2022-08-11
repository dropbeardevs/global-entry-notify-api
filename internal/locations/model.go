package locations

type Locations struct {
	Id                  int    `json:"id" firestore:"id"`
	Name                string `json:"name" firestore:"name"`
	ShortName           string `json:"shortName" firestore:"shortName"`
	LocationType        string `json:"locationType" firestore:"locationType"`
	LocationCode        string `json:"locationCode" firestore:"locationCode"`
	Address             string `json:"address" firestore:"address"`
	AddressAdditional   string `json:"addressAdditional" firestore:"addressAdditional"`
	City                string `json:"city" firestore:"city"`
	State               string `json:"state" firestore:"state"`
	PostalCode          string `json:"postalCode" firestore:"postalCode"`
	CountryCode         string `json:"countryCode" firestore:"countryCode"`
	TimezoneData        string `json:"tzData" firestore:"timezoneData"`
	PhoneNumber         string `json:"phoneNumber" firestore:"phoneNumber"`
	PhoneAreaCode       string `json:"phoneAreaCode" firestore:"phoneAreaCode"`
	PhoneCountryCode    string `json:"phoneCountryCode" firestore:"phoneCountryCode"`
	PhoneExtension      string `json:"phoneExtension" firestore:"phoneExtension"`
	PhoneAltNumber      string `json:"phoneAltNumber" firestore:"phoneAltNumber"`
	PhoneAltAreaCode    string `json:"phoneAltAreaCode" firestore:"phoneAltAreaCode"`
	PhoneAltCountryCode string `json:"phoneAltCountryCode" firestore:"phoneAltCountryCode"`
	PhoneAltExtension   string `json:"phoneAltExtension" firestore:"phoneAltExtension"`
	FaxNumber           string `json:"faxNumber" firestore:"faxNumber"`
	FaxAreaCode         string `json:"faxAreaCode" firestore:"faxAreaCode"`
	FaxCountryCode      string `json:"faxCountryCode" firestore:"faxCountryCode"`
	FaxExtension        string `json:"faxExtension" firestore:"faxExtension"`
	Temporary           string `json:"temporary" firestore:"temporary"`
	InviteOnly          string `json:"inviteOnly" firestore:"inviteOnly"`
	Operational         string `json:"operational" firestore:"operational"`
	Directions          string `json:"directions" firestore:"directions"`
	Notes               string `json:"notes" firestore:"notes"`
	MapFileName         string `json:"mapFileName" firestore:"mapFileName"`
	LastUpdatedDate     string `json:"lastUpdatedDate" firestore:"lastUpdatedDate"`
	Services            []struct {
		Id   int    `json:"id" firestore:"id"`
		Name string `json:"name" firestore:"name"`
	}
}
