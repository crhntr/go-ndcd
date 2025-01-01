package ndcd

import (
	"context"
	"io"
	"net/http"
)

// TextZIP is failing when the request comes from the CLI.
// It seems to work from the browser.
const TextZIP = "http://www.accessdata.fda.gov/cder/ndctext.zip"

// DownloadZip this is not working
func DownloadZip(ctx context.Context, client *http.Client, w io.Writer) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, TextZIP, nil)
	if err != nil {
		return err
	}
	if client == nil {
		client = http.DefaultClient
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(res.Body)
	_, err = io.Copy(w, res.Body)
	return err
}

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}
