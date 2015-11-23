package commands

import (
    "bufio"
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var (
    pushCmd = &cobra.Command{
        Use:   "push [bullet]",
        Short: "Cache new data",
        Long:  "Cache new data if specified key is not used yet or caches (use specified key) are expired.",
        Run:   pushCommand,
    }
)

func pushCommand(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
        cmd.Usage()
        return
    }

    bullet := args[0]
    if bullet == "" {
        Exit(fmt.Errorf("invalid bullet."))
    }
    fr, err := os.OpenFile(getBulletPath(bullet), os.O_RDONLY, 0600)
    if err != nil {
        Exit(err)
    }
    defer fr.Close()

    fw, err := os.OpenFile(getHostsFilePath(), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
    if err != nil {
        Exit(err)
    }
    defer fw.Close()

    err = fw.Truncate(0)
    if err != nil {
        Exit(err)
    }

    scanner := bufio.NewScanner(fr)
    for scanner.Scan() {
        _, err = fw.WriteString(fmt.Sprintln(scanner.Text()))
        if err != nil {
            Exit(err)
        }
    }
    if err = scanner.Err(); err != nil {
        Exit(err)
    }
    fmt.Println("success.")
}

func init() {
    RootCmd.AddCommand(pushCmd)
}