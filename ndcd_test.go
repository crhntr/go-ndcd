package ndcd_test

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

const NDC_ZIP_URL = "http://www.accessdata.fda.gov/cder/ndctext.zip"

func TestMain(m *testing.M) {
	if _, err := os.Stat("testdata"); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		err := os.Mkdir("testdata", 0o700)
		if err != nil {
			panic(err)
		}
	}

	_, packageMissing := os.Stat("testdata/package.txt")
	_, productMissing := os.Stat("testdata/product.txt")

	if packageMissing != nil || productMissing != nil {
		var zipFile *os.File

		if _, zipMissing := os.Stat("testdata/package.txt"); zipMissing != nil {
			getResponse, err := http.Get(NDC_ZIP_URL)
			if err != nil {
				panic(err)
			}
			defer getResponse.Body.Close()

			zipFile, err = os.Create("testdata/ndctext.zip")
			if err != nil {
				log.Panic(err)
			}
			defer zipFile.Close()

			io.Copy(zipFile, getResponse.Body)
		}

		zipReader, err := zip.OpenReader("testdata/ndctext.zip")
		if err != nil {
			log.Panic(err)
		}
		defer zipReader.Close()

		for _, f := range zipReader.File {
			switch f.Name {
			case "product.txt":
				if productMissing == nil {
					continue
				}
				os.Remove("testdata/product.txt")
				file, err := os.Create("testdata/product.txt")
				if err != nil {
					log.Panic(err)
				}
				productFile, err := f.Open()
				if err != nil {
					log.Panic(err)
				}
				io.Copy(file, productFile)
			case "package.txt":
				if packageMissing == nil {
					continue
				}
				os.Remove("testdata/package.txt")
				file, err := os.Create("testdata/package.txt")
				if err != nil {
					log.Panic(err)
				}
				packageFile, err := f.Open()
				if err != nil {
					log.Panic(err)
				}
				io.Copy(file, packageFile)
			}
		}
	}

	flag.Parse()
	os.Exit(m.Run())
}
