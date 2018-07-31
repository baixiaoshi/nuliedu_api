package helper

import (
	"os"
	"errors"
)

//校验目录是否存在
func CheckDirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return fileInfo.IsDir()
}

//创建目录
func CreateDir(dir string) error {
	if !CheckDirExists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return errors.New("create wconf dir fail, dir:" + dir + " err:" + err.Error())
		}
	}
	return nil
}
