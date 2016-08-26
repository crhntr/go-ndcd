package nationaldrugcodedirectory

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const NDC_ZIP_URL = "http://www.accessdata.fda.gov/cder/ndc.zip"

func Download(workingDir string) (string, error) {
	if strings.HasSuffix(workingDir, "/") {
		return "", errors.New("save to path should not end in '/'")
	}
	download(workingDir + "/ndc.zip")
	unzip(workingDir+"/ndc.zip", workingDir+"/product.txt")
	return workingDir + "/product.txt", nil
}

func download(path string) {
	var zipFile *os.File
	var getResponse *http.Response

	getResponse, err := http.Get(NDC_ZIP_URL)
	if err != nil {

		log.Panic(err)
	}

	zipFile, err = os.Create(path)
	if err != nil {
		log.Panic(err)
	}
	defer zipFile.Close()

	io.Copy(zipFile, getResponse.Body)
}

func unzip(from, to string) {
	zipFile, err := zip.OpenReader(from)
	if err != nil {
		log.Panic(err)
	}
	defer zipFile.Close()

	for _, f := range zipFile.File {
		if f.Name == "product.txt" {
			os.Remove(to)
			file, err := os.Create(to)
			if err != nil {
				log.Panic(err)
			}
			productFile, err := f.Open()
			if err != nil {
				log.Panic(err)
			}
			io.Copy(file, productFile)
			os.Remove(from)
		}
	}
}
