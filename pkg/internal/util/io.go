package util

import (
	"os"
	"path/filepath"
)

// CreateDir 检查目录是否存在，不存在则创建，存在则删除后重新创建
func CreateDir(file string) (string, error) {
	absfile, absdir, err := abspath(file)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(absdir); err == nil {
		os.RemoveAll(absdir)
	}

	return absfile, os.MkdirAll(absdir, 0755)
}

// CreateDirIfNotExist 检查目录是否存在，不存在则创建
func CreateDirIfNotExist(file string) (string, error) {
	absfile, absdir, err := abspath(file)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(absdir); os.IsNotExist(err) {
		return absfile, os.MkdirAll(absdir, 0755)
	}

	return absfile, nil
}

// abspath .
func abspath(file string) (absfile string, absdir string, err error) {
	absfile, err = filepath.Abs(file)
	if err != nil {
		return "", "", err
	}

	return absfile, filepath.Dir(absfile), nil
}
