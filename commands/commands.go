package commands

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var (
    RootCmd = &cobra.Command{
        Use:   "holster",
        Short: "manage your hosts file",
        Long:  "manage your hosts file",
        Run:   func(cmd *cobra.Command, args []string) {
            versionCommand(cmd, args)
            cmd.Usage()
        },
    }
    versionCmd = &cobra.Command{
        Use:   "version",
        Short: "Show version",
        Long:  "Show version",
        Run:   versionCommand,
    }
    version = "2.0.0"
)

func versionCommand(cmd *cobra.Command, args []string) {
    fmt.Println(version)
}

func Exit(err error, codes ...int) {
    var code int
    if len(codes) > 0 {
        code = codes[0]
    } else {
        code = 2
    }
    if err != nil {
        fmt.Println(err)
    }
    os.Exit(code)
}

func init() {
    RootCmd.AddCommand(versionCmd)
}

func Run() {
    RootCmd.Execute()
}