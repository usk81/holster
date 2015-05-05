package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    "regexp"
    "strings"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "holster"
    app.Usage = "manage your hosts file"
    app.Version = "0.1.0"
    app.Author = "Yusuke Komatsu"
    app.Commands = []cli.Command{
        {
            Name:        "init",
            Aliases:     []string{"i"},
            Usage:       "initialize hoster's setting",
            Description: "",
            Action:      runInit,
        },
        {
            Name:        "list",
            Aliases:     []string{"l"},
            Usage:       "read hosts file",
            Description: "read hosts file",
            Action:      runReadHosts,
        },
        {
            Name:        "known",
            Aliases:     []string{"k"},
            Usage:       "show hosts group. (list files in .holster/bullets)",
            Description: "show hosts group. (list files in .holster/bullets)",
            Action:      runGetHostsFileList,
        },
        {
            Name:        "update",
            Aliases:     []string{"u"},
            Usage:       "update hosts file",
            Description: "",
            Action:      runUpdate,
        },
        {
            Name:        "append",
            Aliases:     []string{"a"},
            Usage:       "add a host information in hosts file",
            Description: "add a host information in hosts file",
            Action:      runAppend,
        },
        // {
        //     Name:        "bundle",
        //     Aliases:     []string{"b"},
        //     Usage:       "update hosts by bundle setting",
        //     Description: "update hosts by bundle setting",
        //     Action:      runBundle,
        // },
    }

    app.Run(os.Args)
}

func runInit(c *cli.Context) {
    var err error
    err = bootstrap()
    if err != nil {
        fmt.Println("has error in bootstrap action. ", err)
        os.Exit(1)
    }
    err = createSampleBullet()
    if err != nil {
        fmt.Println("has error in create sample bullet. ", err)
        os.Exit(1)
    }
}

func runReadHosts(c *cli.Context) {
    hosts := getHostsFilePath()
    contents, err := ioutil.ReadFile(hosts)
    if err != nil {
        fmt.Println("Can't read hosts file. ", err)
        os.Exit(1)
    }
    fmt.Println(string(contents))
}

func runGetHostsFileList(c *cli.Context) {
    files, err :=  ioutil.ReadDir(filepath.Join(getHomeDir(), ".holster", "bullets"))
    if err != nil {
        fmt.Println("Can't read hosts file. ", err)
        os.Exit(1)
    }
    for _, fi := range files {
        filename := fi.Name()
        if !fi.IsDir() && validateHostsFile(filename) {
            pos := strings.LastIndex(filename, ".")
            println(filename[:pos])
        }
    }
}

func runUpdate(c *cli.Context) {
    bullet := c.Args().First()
    if bullet == "" {
        fmt.Println("invalid bullet.")
        os.Exit(1)
    }
    err := update(bullet)
    if err != nil {
        fmt.Println("error in updating hosts file. ", err)
        os.Exit(1)
    }
}

func runAppend(c *cli.Context) {
    ip   := c.Args().Get(0)
    if ip == "" {
        fmt.Println("invalid ipaddress.")
        os.Exit(1)
    }
    host := c.Args().Get(1)
    if host == "" {
        fmt.Println("invalid host.")
        os.Exit(1)
    }
    err := append(ip, host)
    if err != nil {
        fmt.Println("error in apppending hosts file. ", err)
        os.Exit(1)
    }
}

// func runBundle(c *cli.Context) {

// }

func update(bullet string) error {
    fr, err := os.OpenFile(getBulletPath(bullet), os.O_RDONLY, 0600)
    if err != nil {
        return err
    }
    defer fr.Close()

    fw, err := os.OpenFile(getHostsFilePath(), os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer fw.Close()
    err = fw.Truncate(0)
    if err != nil {
        return err
    }

    scanner := bufio.NewScanner(fr)
    for scanner.Scan() {
        _, err = fw.WriteString(fmt.Sprintln(scanner.Text()))
        if err != nil {
            return err
        }
    }
    if err = scanner.Err(); err != nil {
        return err
    }
    return nil
}

func append(ipaddr string, hosts string) error {
    file, err := os.OpenFile(getHostsFilePath(), os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    line := fmt.Sprintf("%s\t%s", ipaddr, hosts)
    _, err = file.WriteString(line)
    if err != nil {
        return err
    }
    return nil
}

func validateHostsFile(filename string) bool {
    match, _ := regexp.MatchString("\\.host$", filename)
    return match
}