package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadJSON will load data from a JSON file into application and return a bytes array
func LoadJSON(path string) ([]byte, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return []byte(``), err
	}

	jsonFile, err := os.Open(p)
	defer jsonFile.Close()
	if err != nil {
		return []byte(``), err
	}

	return ioutil.ReadAll(jsonFile)
}
