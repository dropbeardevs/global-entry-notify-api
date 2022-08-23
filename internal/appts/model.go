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
}

type wsAppointment struct {
	LocationId     int    `json:"locationId"`
	StartTimestamp string `json:"startTimestamp"`
}
type dbAppointment struct {
	LocationId        int             `bson:"locationId"`
	Date              time.Time       `bson:"date"`
	NotificationsList []Notifications `bson:"notificationsList"`
	LastUpdated       time.Time       `bson:"lastUpdated"`
}
