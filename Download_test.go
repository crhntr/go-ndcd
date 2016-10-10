package nationaldrugcodedirectory_test

import (
	"testing"

	ndcd "github.com/crhntr/nationaldrugcodedirectory"
)

func TestDownload(t *testing.T) {
	t.SkipNow()
	if _, err := ndcd.Download("testdata/"); err == nil {
		t.Error("should not allow path ending in '/'")
	}

	path, err := ndcd.Download("testdata")
	if err != nil {
		t.Error(err)
	}
	if path == "" {
		t.Error("empty path but no error")
	}
}
