package commands

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var (
    removeCmd = &cobra.Command{
        Use:   "remove [bullet]",
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
        Run:   removeCommand,
    }
)

func removeCommand(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
        bullet := args[0]
        if bullet == "" {
            Exit(fmt.Errorf("invalid bullet."))
        }
        fp := getBulletPath(bullet)
        if !existFile(fp) {
            Exit(fmt.Errorf("this bullet is not exist."))
        }

        if !askForConfirmation(fmt.Sprintf("remove %s, really? (y/n)", bullet)) {
            Exit(nil, 1)
        }

        if err := os.Remove(fp); err != nil {
            Exit(err)
        }
    } else {
        cmd.Usage()
    }
}

func init() {
    RootCmd.AddCommand(removeCmd)
}