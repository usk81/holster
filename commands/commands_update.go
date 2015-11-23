package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
    updateCmd = &cobra.Command{
        Use:   "update [bullet]",
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
        Run:   updateCommand,
    }
)

func updateCommand(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
        bullet := args[0]
        if bullet == "" {
            Exit(fmt.Errorf("invalid bullet."))
        }
        fp := getBulletPath(bullet)
        if err := editFile(fp); err != nil {
            fmt.Printf("Error while editing. Error: %v\n", err)
        } else {
            fmt.Println("Update success.")
        }
    } else {
        cmd.Usage()
    }
}

func init() {
    RootCmd.AddCommand(updateCmd)
}