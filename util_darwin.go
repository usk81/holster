// +build darwin
package main

import (
	"os"
    "path/filepath"
)

func getHomeDir() string {
    home := os.Getenv("HOLSTER_HOME")
    if home != "" {
        home = os.Getenv("HOME")
    }
    return home
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
    file := filepath.Join("private", "etc", "hosts")
    file = string(os.PathSeparator) + file
    return file
}

func getBulletPath(bullet string) string {
    bulletFile := bullet + ".host"
    return filepath.Join(getHomeDir(), ".holster", "bullets", bulletFile)
}