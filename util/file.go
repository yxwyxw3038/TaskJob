package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetBasePath() (string, error) {
	baseDirPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return baseDirPath, nil
}
func GetFileInfo(fileName string) (string, error) {
	baseDirPath, err := GetBasePath()
	if err != nil {

		return "", err
	}
	filePath := baseDirPath + "\\tasklist\\" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	msg := string(b)
	return msg, nil
}
