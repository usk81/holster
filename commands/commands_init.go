package commands

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "github.com/spf13/cobra"
)

var (
    initCmd = &cobra.Command{
        Use:   "init",
        Short: "initialize hoster's setting",
        Long:  "initialize hoster's setting",
        Run:   initCommand,
    }
)

func initCommand(cmd *cobra.Command, args []string) {
    var err error
    err = bootstrap()
    if err != nil {
        Exit(fmt.Errorf("has error in bootstrap action. %v", err))
    }
    err = createSampleBullet()
    if err != nil {
        Exit(fmt.Errorf("has error in create sample bullet. %v", err))
    }
}

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
    file := filepath.Join(home, ".holster", "bullets", "default.host")
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
func init() {
    RootCmd.AddCommand(initCmd)
}