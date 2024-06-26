package util

import (
	"os"
	"path/filepath"
)

// CreateDir 检查文件所在目录是否存在，如果不存在则创建它，并返回文件全路径
func CreateDir(f string, recreate bool) (string, error) {
	absfile, err := filepath.Abs(f)
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(absfile)

	_, err = os.Stat(dir)
	if os.IsExist(err) {
		// 不需要重新创建
		if !recreate {
			return absfile, nil
		}

		os.RemoveAll(dir)
	}

	return absfile, os.MkdirAll(dir, 0755)
}
