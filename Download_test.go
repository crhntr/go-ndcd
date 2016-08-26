package nationaldrugcodedirectory_test

import (
	"testing"

	ndcd "github.com/hunteramericano/nationaldrugcodedirectory"
)

func TestDownload(t *testing.T) {
	NationalDrugCodeDirectory.Parse(ndcd.GetLatest(), 100)
}
