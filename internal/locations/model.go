package locations

import "time"

type Locations struct {
	LocationId              int    `json:"id" bson:"locationId"`
	Name                    string `json:"name" bson:"name"`
	ShortName               string `json:"shortName" bson:"shortName"`
	LocationType            string `json:"locationType" bson:"locationType"`
	LocationCode            string `json:"locationCode" bson:"locationCode"`
	Address                 string `json:"address" bson:"address"`
	AddressAdditional       string `json:"addressAdditional" bson:"addressAdditional"`
	City                    string `json:"city" bson:"city"`
	State                   string `json:"state" bson:"state"`
	PostalCode              string `json:"postalCode" bson:"postalCode"`
	CountryCode             string `json:"countryCode" bson:"countryCode"`
	TimezoneData            string `json:"tzData" bson:"timezoneData"`
	PhoneNumber             string `json:"phoneNumber" bson:"phoneNumber"`
	PhoneAreaCode           string `json:"phoneAreaCode" bson:"phoneAreaCode"`
	PhoneCountryCode        string `json:"phoneCountryCode" bson:"phoneCountryCode"`
	PhoneExtension          string `json:"phoneExtension" bson:"phoneExtension"`
	PhoneAltNumber          string `json:"phoneAltNumber" bson:"phoneAltNumber"`
	PhoneAltAreaCode        string `json:"phoneAltAreaCode" bson:"phoneAltAreaCode"`
	PhoneAltCountryCode     string `json:"phoneAltCountryCode" bson:"phoneAltCountryCode"`
	PhoneAltExtension       string `json:"phoneAltExtension" bson:"phoneAltExtension"`
	FaxNumber               string `json:"faxNumber" bson:"faxNumber"`
	FaxAreaCode             string `json:"faxAreaCode" bson:"faxAreaCode"`
	FaxCountryCode          string `json:"faxCountryCode" bson:"faxCountryCode"`
	FaxExtension            string `json:"faxExtension" bson:"faxExtension"`
	Temporary               bool   `json:"temporary" bson:"temporary"`
	InviteOnly              bool   `json:"inviteOnly" bson:"inviteOnly"`
	Operational             bool   `json:"operational" bson:"operational"`
	Directions              string `json:"directions" bson:"directions"`
	Notes                   string `json:"notes" bson:"notes"`
	MapFileName             string `json:"mapFileName" bson:"mapFileName"`
	LocationLastUpdatedDate string `json:"lastUpdatedDate" bson:"locationLastUpdatedDate"`
	Services                []struct {
		Id   int    `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	}
	LastUpdated time.Time `bson:"lastUpdated"`
}
