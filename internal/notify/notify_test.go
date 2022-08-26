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

func TestUpdateNotificationDetails(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()
	var err error

	userId := "3D05A979-35F9-4A40-B075-444DEB63537A"
	targetDate, _ := time.Parse("2006-01-02", "2023-12-31")

	notificationDetails := models.NotificationDetails{LocationId: 5001, TargetDate: targetDate}

	notificationDetails2 := models.NotificationDetails{LocationId: 5180, TargetDate: targetDate}

	notificationDetails3 := models.NotificationDetails{LocationId: 16461, TargetDate: targetDate}

	// Execute function
	err = UpdateNotificationDetails(userId, notificationDetails)
	if err != nil {
		t.Fatalf("UpdateNotificationDetails(%v) = %v, want match for nil\n", userId, err)
	} else {
		sugar.Infof("UpdateNotificationDetails(%v) Success\n", userId)
	}

	err = UpdateNotificationDetails(userId, notificationDetails2)
	if err != nil {
		t.Fatalf("UpdateNotificationDetails(%v) = %v, want match for nil\n", userId, err)
	} else {
		sugar.Infof("UpdateNotificationDetails(%v) Success\n", userId)
	}

	err = UpdateNotificationDetails(userId, notificationDetails3)
	if err != nil {
		t.Fatalf("UpdateNotificationDetails(%v) = %v, want match for nil\n", userId, err)
	} else {
		sugar.Infof("UpdateNotificationDetails(%v) Success\n", userId)
	}

}

func TestDeleteNotificationDetails(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()
	var err error

	userId := "3D05A979-35F9-4A40-B075-444DEB63537A"
	locationId := 5001

	// Execute function
	err = DeleteNotificationDetails(userId, locationId)
	if err != nil {
		t.Fatalf("DeleteNotificationDetails(%v, %v) = %v, want match for nil\n", userId, locationId, err)
	} else {
		sugar.Infof("DeleteNotificationDetails(%v, %v) Success\n", userId, locationId)
	}
}

func TestUpdateNotificationToken(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()
	var err error

	userId1 := "3D05A979-35F9-4A40-B075-444DEB63537A"
	userId2 := "567D334A-62FA-456C-9065-6E1DD7FED0A1"
	userId3 := "D698CD9B-4C73-4168-BBB2-FDA3CCB10C40"

	// Execute function
	err = UpdateNotificationToken(userId1, uuid.NewString())
	if err != nil {
		t.Fatalf("UpdateNotificationToken(%v) = %v, want nil", userId1, err)
	} else {
		sugar.Infof("UpdateNotification(%v) Success\n", userId1)
	}

	err = UpdateNotificationToken(userId2, uuid.NewString())
	if err != nil {
		t.Fatalf("UpdateNotificationToken(%v) = %v, want nil", userId2, err)
	} else {
		sugar.Infof("UpdateNotification(%v) Success\n", userId2)
	}

	err = UpdateNotificationToken(userId3, uuid.NewString())
	if err != nil {
		t.Fatalf("UpdateNotificationToken(%v) = %v, want nil", userId3, err)
	} else {
		sugar.Infof("UpdateNotification(%v) Success\n", userId3)
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
