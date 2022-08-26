package notify

import (
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/initapp"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"github.com/google/uuid"
)

func init() {
	// Get current running filename
	_, filename, _, _ := runtime.Caller(0)

	// Change to ../..
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	initapp.InitApp()

	sugar := logger.GetInstance()

	sugar.Infof("Running folder: %v", dir)
}

func TestAddNotificationDetails(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()
	want := true
	date, _ := time.Parse("2006-01-02", "2022-09-31")
	var err error

	notification := models.Notification{
		UserId:      "3D05A979-35F9-4A40-B075-444DEB63537A",
		LocationIds: []int{5001, 5002, 5003},
		Token:       uuid.NewString(),
		TargetDate:  date,
	}

	notification2 := models.Notification{
		UserId:      "567D334A-62FA-456C-9065-6E1DD7FED0A1",
		LocationIds: []int{5180},
		Token:       uuid.NewString(),
		TargetDate:  date,
	}

	notification3 := models.Notification{
		UserId:      "D698CD9B-4C73-4168-BBB2-FDA3CCB10C40",
		LocationIds: []int{16461, 8120, 5446},
		Token:       uuid.NewString(),
		TargetDate:  date,
	}

	// Execute function
	err = UpdateNotification(notification)
	if err != nil {
		t.Fatalf("UpdateNotification(%v) = %v, want match for %t, nil\n", notification, err, want)
	} else {
		sugar.Infof("UpdateNotification(%v) Success\n", notification)
	}

	err = UpdateNotification(notification2)
	if err != nil {
		t.Fatalf("UpdateNotification(%v) = %v, want match for %t, nil\n", notification2, err, want)
	} else {
		sugar.Infof("UpdateNotification(%v) Success\n", notification)
	}

	err = UpdateNotification(notification3)
	if err != nil {
		t.Fatalf("UpdateNotification(%v) = %v, want match for %t, nil\n", notification3, err, want)
	} else {
		sugar.Infof("UpdateNotification(%v) Success\n", notification)
	}

}

func TestGetDbNotifications(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()

	// Execute function
	err := getDbNotifications()

	if err != nil {
		t.Fatalf("getDbNotifications() = %v", err)
	} else {
		sugar.Infof("getDbNotifications() Success\n")
	}

}
