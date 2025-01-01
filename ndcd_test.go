package ndcd_test

import (
	"context"
	"io"
	"testing"

	"github.com/crhntr/go-ndcd"
)

func TestDownloadZip(t *testing.T) {
	t.Skip("the server thinks this request is abuse")
	ctx := context.TODO()
	if err := ndcd.DownloadZip(ctx, nil, io.Discard); err != nil {
		t.Fatal(err)
	}
}
