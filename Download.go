package nationaldrugcodedirectory

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
)

const pNDC_ZIP_URL = "http://www.accessdata.fda.gov/cder/ndc.zip"
const pNDC_ZIP_PATH = "temp/ndc.zip"
const pNDC_PRODUCT_TXT_PATH = "temp/product.txt"

func GetLatest(logger *log.Logger) string {
	download(logger)
	unzip(logger)
	return pNDC_PRODUCT_TXT_PATH
}

func download(logger *log.Logger) {
	var zipFile *os.File
	var getResponse *http.Response

	logger.Printf("Downloading ndc.zip from: %s", pNDC_ZIP_URL)
	getResponse, err := http.Get(pNDC_ZIP_URL)
	if err != nil {
		logger.Panic(err)
	}
	logger.Print("Finished Downloading ndc.zip")

	logger.Printf("Done : %s", pNDC_ZIP_URL)
	zipFile, err = os.Create(pNDC_ZIP_PATH)
	if err != nil {
		logger.Panic(err)
	}
	defer func() {
		zipFile.Close()
		if err != nil {
			logger.Panic(err)
		}
	}()

	logger.Printf("Saving ndc.zip to: %s", pNDC_ZIP_PATH)
	io.Copy(zipFile, getResponse.Body)
}

func unzip(logger *log.Logger) {
	zipFile, err := zip.OpenReader(pNDC_ZIP_PATH)
	if err != nil {
		logger.Panic(err)
	}
	defer zipFile.Close()

	for _, f := range zipFile.File {
		if f.Name == "product.txt" {
			os.Remove(pNDC_PRODUCT_TXT_PATH)
			file, err := os.Create(pNDC_PRODUCT_TXT_PATH)
			if err != nil {
				logger.Panic(err)
			}
			productFile, err := f.Open()
			if err != nil {
				logger.Panic(err)
			}
			io.Copy(file, productFile)
			os.Remove(pNDC_ZIP_PATH)
		}
	}
	logger.Printf("Unzipped and saved product.txt to: %s", pNDC_ZIP_PATH)
}
