package locations

import (
	"os"
	"path"
	"runtime"
	"testing"

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

	sugar := logger.GetInstance()

	sugar.Infof("Running folder: %v", dir)
}

func TestGetLocations(t *testing.T) {

	// Bootstrap code
	sugar := logger.GetInstance()
	var err error

	// Execute function
	locationList, err := GetLocations()
	if err != nil {
		t.Fatalf("GetLocations() = %v, want match for nil\n", err)
	} else if len(*locationList) == 0 {
		t.Fatalf("GetLocations() = %v, no locations found\n", locationList)
	} else {
		sugar.Infof("GetLocations() Success: %#v", locationList)
	}

}
