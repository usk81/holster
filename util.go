package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "path/filepath"
)

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

func getBulletPath(bullet string) string {
    bulletFile := bullet + ".host"
    return filepath.Join(getHomeDir(), ".holster", "bullets", bulletFile)
}

func editFile(fp string) error {
    if runtime.GOOS == "windows" {
        return fmt.Errorf("not support to edit files.")
    }

    cmd := exec.Command("vim", fp)
    cmd.Stdin  = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    var err error
    if err = cmd.Start(); err != nil {
        return err
    }

    err = cmd.Wait()
    if err != nil {
        return err
    }
    return nil
}
