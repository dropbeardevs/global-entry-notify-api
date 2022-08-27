package appts

import (
	"os"
	"path"
	"runtime"
	"testing"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/initapp"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
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

func TestGetDbAppointments(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()

	// Execute function
	err := getDbAppointments()
	if err != nil {
		t.Fatalf("getDbAppointments() = %v, want match for nil\n", err)
	} else {
		sugar.Infof("getDbAppointments() Success\n")
	}
}
