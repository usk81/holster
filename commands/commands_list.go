package commands

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"
    "github.com/spf13/cobra"
)

var (
    listCmd = &cobra.Command{
        Use:   "list",
        Short: "show hosts group",
        Long:  "show hosts group. (list files in .holster/bullets)",
        Run:   listCommand,
    }
)

func listCommand(cmd *cobra.Command, args []string) {
    files, err :=  ioutil.ReadDir(filepath.Join(getHomeDir(), ".holster", "bullets"))
    if err != nil {
        Exit(fmt.Errorf("Can't read hosts file. %v", err))
    }
    for _, fi := range files {
        filename := fi.Name()
        if !fi.IsDir() && validateHostsFile(filename) {
            pos := strings.LastIndex(filename, ".")
            println(filename[:pos])
        }
    }
}

func init() {
    RootCmd.AddCommand(listCmd)
}