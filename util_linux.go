// +build linux
package main

import (
	"os"
    "path/filepath"
)

func getHomeDir() string {
	home := os.Getenv("HOME")
	return home
}

func existFile(file_path string) bool {
    src, err := os.Stat(file_path)
    if err != nil {
        return false
    }

    if !src.IsDir() {
        return true
    }
    return false
}

func existDirectory(dir_path string) bool {
	src, err := os.Stat(dir_path)
	if err != nil {
		return false
	}

	if !src.IsDir() {
		return false
	}
	return true
}

func getHostsFilePath() string {
    file := filepath.Join("etc", "hosts")
    file = string(os.PathSeparator) + file
    return file
}

func getBulletPath(bullet string) string {
    bulletFile := bullet + ".host"
    return filepath.Join(getHomeDir(), ".holster", "bullets", bulletFile)
}