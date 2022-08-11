package appts

type ApptUrlValues struct {
	LocationId int
}

type Appointment struct {
	LocationId int    `json:"locationId"`
	StartDate  string `json:"startTimestamp"`
	Active     bool   `json:"active"`
}
