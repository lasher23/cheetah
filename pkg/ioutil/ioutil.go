package ioutil

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func CreateEmptyFileWithAllDirectories(location string, filename string) error {
	os.MkdirAll(location, os.ModeDir)
	file, e := os.Create(location + "\\" + filename)
	defer file.Close()
	return e
}

func ReadFile(location string) (string, error) {
	bytes, e := ioutil.ReadFile(location)
	if e != nil {
		return "", errors.Wrap(e, "error accessing file")
	}
	return string(bytes), nil
}
