package notify

import (
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
)

func init() {

	homeDir, _ := os.UserHomeDir()
	os.Setenv("GLOBAL-ENTRY-NOTIFY-API-CONFIG", homeDir+"/Developer/secrets/global-entry-notify-api/config.yml")
	defer os.Unsetenv("GLOBAL-ENTRY-NOTIFY-API-CONFIG")

	// Get current running filename
	_, filename, _, _ := runtime.Caller(0)

	// Change to ../..
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	sugar := logger.GetInstance()

	sugar.Infof("Running folder: %v", dir)
}

func TestUpdateNotificationDetails(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()
	var err error

	userId := "3D05A979-35F9-4A40-B075-444DEB63537A"
	targetDate, _ := time.Parse("2006-01-02", "2022-10-31")

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

	userId := "3D05A979-35F9-4A40-B075-444DEB63537A"
	token := "fd8faYbLTSOA-pJzIKEKbp:APA91bEvXKtDsFu1Uowv5Ubeg6ZNCSr3fxPfp6R1PmJ7YoHrwz5O1meFqdt1y2g82W1dzNkqAwGF5R6hL--YxBQK421SaslDl0BGGLcGbw2rWNhkJtw9e-upR2xibQ29ckjvX837cAQ3"

	// Execute function
	err = UpdateNotificationToken(userId, token)
	if err != nil {
		t.Fatalf("UpdateNotificationToken(%v, %v) = %v, want nil", userId, token, err)
	} else {
		sugar.Infof("UpdateNotification(%v, %v) Success\n", userId, token)
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

func TestGetNotificationDetails(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()

	userId := "3D05A979-35F9-4A40-B075-444DEB63537A"

	// Execute function
	result, err := GetNotificationDetails(userId)

	if err != nil {
		t.Fatalf("GetNotificationDetails(%v) = %v", userId, err)
	} else {
		sugar.Infof("GetNotificationDetails(%v) result: %#v", userId, result)
	}

}

func TestSendNotification(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()

	token := "9c108294-7c08-4939-b697-6385ee93e376"
	//token := "fd8faYbLTSOA-pJzIKEKbp:APA91bEvXKtDsFu1Uowv5Ubeg6ZNCSr3fxPfp6R1PmJ7YoHrwz5O1meFqdt1y2g82W1dzNkqAwGF5R6hL--YxBQK421SaslDl0BGGLcGbw2rWNhkJtw9e-upR2xibQ29ckjvX837cAQ3"

	dir, _ := os.UserHomeDir()

	dir = dir + "/Developer/secrets/global-entry-notify-api/global-entry-c8373-fe72e1ae9c11.json"

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", dir)
	defer os.Unsetenv("ENV_VAR")

	time := time.Now()
	dateLayout := "January 2, 2006 3:04 PM"

	msg := "New Global Entry Appointment Available at " + time.Format(dateLayout)

	// Execute function
	result, err := sendNotification(token, msg)

	if err != nil {
		t.Fatalf("getDbNotifications() = %v", err)
	} else {
		sugar.Infof("getDbNotifications() Success, result: %v", result)
	}

}
