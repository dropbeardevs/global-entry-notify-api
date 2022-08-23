package appts

import "time"

type ApptUrlValues struct {
	LocationId int
}

type Notifications struct {
	UserId           string    `bson:"userId"`     // UUID
	Token            string    `bson:"token"`      // Firebase Cloud Messaging Token
	TargetDate       time.Time `bson:"targetDate"` // Want to get notifications for appointments before this date
	LastNotifiedDate time.Time `bson:"lastNotifiedDate"`
	LastUpdated      time.Time `bson:"lastUpdated"`
}

type Appointment struct {
	LocationId        int             `json:"locationId" bson:"locationId"`
	StartDate         string          `json:"startTimestamp" bson:"startDate"`
	Active            bool            `json:"active" bson:"active"`
	NotificationsList []Notifications `bson:"notificationsList"`
	LastUpdated       time.Time       `bson:"lastUpdated"`
}
