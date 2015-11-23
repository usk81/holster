package commands

import (
    "fmt"
    "io/ioutil"
    "github.com/spf13/cobra"
)

var (
    currentCmd = &cobra.Command{
        Use:   "current",
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
        Run:   currentCommand,
    }
)

func currentCommand(cmd *cobra.Command, args []string) {
    hosts := getHostsFilePath()
    contents, err := ioutil.ReadFile(hosts)
    if err != nil {
        Exit(fmt.Errorf("Can't read hosts file. ", err))
    }
    fmt.Println(string(contents))
}

func init() {
    RootCmd.AddCommand(currentCmd)
}