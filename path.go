package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetPathDirs 获取PATH 所设置的所有文件夹
func GetPathDirs() []string {
	paths := os.Getenv("PATH")
	if runtime.GOOS == "windows" {
		return strings.Split(paths, ";")
	}
	return strings.Split(paths, ":")
}

// ListDir list file in dir
func ListDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, f := range files {
		list = append(list, filepath.Join(dir, f.Name()))
	}
	return list, nil
}
