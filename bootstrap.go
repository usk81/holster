package main

import (
    "io/ioutil"
    "path/filepath"
    "os"
)

func bootstrap() error {
    var err error
    home := getHomeDir()

    // Exists .holster directory
    holsterDir := filepath.Join(home, ".holster")
    if !existDirectory(holsterDir) {
        err = os.Mkdir(holsterDir, 0777)
        if err != nil {
            return err
        }
    }

    // Exists .holster/bullets directory
    bulletsDir := filepath.Join(holsterDir, "bullets")
    if !existDirectory(bulletsDir) {
        err = os.Mkdir(bulletsDir, 0777)
        if err != nil {
            return err
        }
    }

    // Exists .holster/bundles directory
    bundlesDir := filepath.Join(holsterDir, "bundles")
    if !existDirectory(bundlesDir) {
        err = os.Mkdir(bundlesDir, 0777)
        if err != nil {
            return err
        }
    }

    return nil   
}

func createSampleBullet() error {
    var err error
    home := getHomeDir()
    file := filepath.Join(home, ".holster", "bullets", "sample.host")
    _, err = os.Stat(file)
    if err != nil {
        host := []byte(
            "##\n" +
            "# Host Database\n" +
            "#\n" +
            "# localhost is used to configure the loopback interface\n" +
            "# when the system is booting.  Do not change this entry.\n" +
            "##\n" +
            "127.0.0.1       localhost\n" +
            "255.255.255.255 broadcasthost\n" +
            "::1             localhost\n",
        )
        err = ioutil.WriteFile(file, host, 0644)
        if err != nil {
            return err
        }
        return nil
    }
    return nil
}