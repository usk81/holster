package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
    addCmd = &cobra.Command{
        Use:   "set [key] [value] [expire]",
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
        Run:   addCommand,
    }
)

func addCommand(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
        bullet := args[0]
        if bullet == "" {
            Exit(fmt.Errorf("invalid bullet."))
        }
        fp := getBulletPath(bullet)
        if existFile(fp) == true {
            Exit(fmt.Errorf("this bullet is not exist."))
        }
        if err := editFile(fp); err != nil {
            Exit(fmt.Errorf("Error while editing. Error: %v", err))
        } else {
            fmt.Println("Create success.")
        }
    } else {
        cmd.Usage()
    }
}

func init() {
    RootCmd.AddCommand(addCmd)
}