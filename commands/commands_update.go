package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
    updateCmd = &cobra.Command{
        Use:   "update [bullet]",
        Short: "update a bullet file",
        Long:  "update a bullet file (preset host file in .holster/bullets)",
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