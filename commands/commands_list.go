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
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
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