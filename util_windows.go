// +build windows
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func getHomeDir() string {
    drive := os.Getenv("HOMEDRIVE")
    path  := os.Getenv("HOMEPATH")
    home  := drive + path
    if drive == "" || path == "" {
        home = os.Getenv("USERPROFILE")
    }
    if home == "" {
        fmt.Println("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
        return ""
    }
    return home
}

func getHostsFilePath() string {
    sysRoot := os.Getenv("SystemRoot")
    if sysRoot == "" {
        fmt.Println("SystemRoot are blank")
        return ""
    }

    file := filepath.Join(sysRoot, "System32", "drivers", "etc", "hosts")
    return file
}