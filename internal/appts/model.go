package appts

import "time"

type ApptUrlValues struct {
	LocationId int
}

type Notifications struct {
	Token        string    `bson:"token"`
	TargetDate   time.Time `bson:"targetDate"` // Want to get notifications for appointments before this date
	LastNotified time.Time `bson:"lastNotified"`
	LastUpdated  time.Time `bson:"lastUpdated"`
}

type Appointment struct {
	LocationId        int             `json:"locationId" bson:"locationId"`
	StartDate         string          `json:"startTimestamp" bson:"startDate"`
	Active            bool            `json:"active" bson:"active"`
	NotificationsList []Notifications `bson:"notificationsList"`
	LastUpdated       time.Time       `bson:"lastUpdated"`
}
