package notify

import (
	"log"
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
	want := true
	date, _ := time.Parse("2006-01-02", "2022-09-31")
	var err error

	notification := models.NotificationDetails{
		LocationIds: []int{5001, 5002, 5003},
		Token:       uuid.NewString(),
		TargetDate:  date,
	}

	// Execute function
	err = AddNotificationDetails(notification)

	if err != nil {
		t.Fatalf("AddNotificationDetails(%v) = %v, want match for %t, nil\n", notification, err, want)
	} else {
		log.Printf("AddNotificationDetails(%v) Success\n", notification)
	}

}
