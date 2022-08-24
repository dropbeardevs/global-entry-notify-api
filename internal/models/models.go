package models

import "time"

type ApptUrlValue struct {
	LocationId int
}

type Notification struct {
	Token            string    `bson:"token"`      // Firebase Cloud Messaging Token
	TargetDate       time.Time `bson:"targetDate"` // Want to get notifications for appointments before this date
	LastNotifiedDate time.Time `bson:"lastNotifiedDate"`
}

type NotificationDetails struct {
	LocationIds []int     `bson:"locationIds"` // Selected Location IDs
	Token       string    `bson:"token"`       // Firebase Cloud Messaging Token
	TargetDate  time.Time `bson:"targetDate"`  // Want to get notifications for appointments before this date
}

type WsAppointment struct {
	LocationId     int    `json:"locationId"`
	StartTimestamp string `json:"startTimestamp"`
}
type DbAppointment struct {
	LocationId       int            `bson:"locationId"`
	Date             time.Time      `bson:"date"`
	NotificationList []Notification `bson:"notificationsList"`
	LastUpdated      time.Time      `bson:"lastUpdated"`
}

type Location struct {
	LocationId            int    `json:"id" bson:"locationId"`
	Name                  string `json:"name" bson:"name"`
	ShortName             string `json:"shortName" bson:"shortName"`
	LocationType          string `json:"locationType" bson:"locationType"`
	LocationCode          string `json:"locationCode" bson:"locationCode"`
	Address               string `json:"address" bson:"address"`
	AddressAdditional     string `json:"addressAdditional" bson:"addressAdditional"`
	City                  string `json:"city" bson:"city"`
	State                 string `json:"state" bson:"state"`
	PostalCode            string `json:"postalCode" bson:"postalCode"`
	CountryCode           string `json:"countryCode" bson:"countryCode"`
	TimezoneData          string `json:"tzData" bson:"timezoneData"`
	PhoneNumber           string `json:"phoneNumber" bson:"phoneNumber"`
	PhoneAreaCode         string `json:"phoneAreaCode" bson:"phoneAreaCode"`
	PhoneCountryCode      string `json:"phoneCountryCode" bson:"phoneCountryCode"`
	PhoneExtension        string `json:"phoneExtension" bson:"phoneExtension"`
	FaxNumber             string `json:"faxNumber" bson:"faxNumber"`
	FaxAreaCode           string `json:"faxAreaCode" bson:"faxAreaCode"`
	FaxCountryCode        string `json:"faxCountryCode" bson:"faxCountryCode"`
	FaxExtension          string `json:"faxExtension" bson:"faxExtension"`
	Temporary             bool   `json:"temporary" bson:"temporary"`
	InviteOnly            bool   `json:"inviteOnly" bson:"inviteOnly"`
	Operational           bool   `json:"operational" bson:"operational"`
	Directions            string `json:"directions" bson:"directions"`
	Notes                 string `json:"notes" bson:"notes"`
	MapFileName           string `json:"mapFileName" bson:"mapFileName"`
	WsLocationLastUpdated string `json:"lastUpdatedDate" bson:"wsLocationLastUpdated"`
	Services              []struct {
		Id   int    `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	}
	LastUpdated time.Time `bson:"lastUpdated"`
}

type Creds struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

type Configuration struct {
	Urls                       map[string]string `yaml:"urls"`
	ConnectionString           string            `yaml:"connectionString"`
	AppointmentPollingTime     int               `yaml:"appointmentPollingTime"`
	LocationsPollingTime       int               `yaml:"locationsPollingTime"`
	AppointmentListPollingTime int               `yaml:"appointmentListPollingTime"`
	LogFileLocation            string            `yaml:"logFileLocation"`
	ZapDefaultLogLevel         string            `yaml:"zapDefaultLogLevel"`
}
