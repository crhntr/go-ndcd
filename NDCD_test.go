package ndcd_test

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if _, err := os.Stat("testdata"); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		err := os.Mkdir("testdata", 0700)
		if err != nil {
			panic(err)
		}
	}
	flag.Parse()
	os.Exit(m.Run())
}
