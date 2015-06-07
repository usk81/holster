// +build linux
package main

import (
	"os"
    "path/filepath"
)

func getHomeDir() string {
    home := os.Getenv("HOLSTER_HOME")
    if home == "" {
        home = os.Getenv("HOME")
    }
	return home
}

func getHostsFilePath() string {
    file := filepath.Join("etc", "hosts")
    file = string(os.PathSeparator) + file
    return file
}
