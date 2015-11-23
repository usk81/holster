package commands

import (
    "fmt"
    "io/ioutil"
    "github.com/spf13/cobra"
)

var (
    showCmd = &cobra.Command{
        Use:   "show [bullet]",
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
        Run:   showCommand,
    }
)

func showCommand(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
        bullet := args[0]
        if bullet == "" {
            Exit(fmt.Errorf("invalid bullet."))
        }
        fp := getBulletPath(bullet)
        if !existFile(fp) {
            Exit(fmt.Errorf("this bullet is not exist."))
        }

        contents, err := ioutil.ReadFile(fp)
        if err != nil {
            Exit(fmt.Errorf("Can't read bullet file. %v", err))
        }
        fmt.Println(string(contents))
    } else {
        cmd.Usage()
    }
}

func init() {
    RootCmd.AddCommand(showCmd)
}