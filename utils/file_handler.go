package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

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
