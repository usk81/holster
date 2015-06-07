package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "github.com/codegangsta/cli"
)

func bulletAdd(c *cli.Context) {
    bullet := c.Args().First()
    if bullet == "" {
        fmt.Println("invalid bullet.")
        os.Exit(1)
    }

    fp := getBulletPath(bullet)
    if existFile(fp) == true {
        fmt.Println("this bullet is not exist.")
        os.Exit(1)
    }

    if err := editFile(fp); err != nil {
        fmt.Printf("Error while editing. Error: %v\n", err)
    } else {
        fmt.Println("Create success.")
    }
}

func bulletUpdate(c *cli.Context) {
    bullet := c.Args().First()
    if bullet == "" {
        fmt.Println("invalid bullet.")
        os.Exit(1)
    }

    fp := getBulletPath(bullet)

    // vim
    if err := editFile(fp); err != nil {
        fmt.Printf("Error while editing. Error: %v\n", err)
    } else {
        fmt.Println("Update success.")
    }
}

func bulletRemove(c *cli.Context) {
    bullet := c.Args().First()
    if bullet == "" {
        fmt.Println("invalid bullet.")
        os.Exit(1)
    }

    fp := getBulletPath(bullet)
    if existFile(fp) == false {
        fmt.Println("this bullet is not exist.")
        os.Exit(1)
    }

    // interactive interface
    if askForConfirmation(fmt.Sprintf("remove %s, really? (y/n)", bullet)) == false {
        os.Exit(1)
    }

    if err := os.Remove(fp); err != nil {
        fmt.Println(err)
    }
}

func bulletShow(c *cli.Context) {
    bullet := c.Args().First()
    if bullet == "" {
        fmt.Println("invalid bullet.")
        os.Exit(1)
    }

    fp := getBulletPath(bullet)
    if existFile(fp) == false {
        fmt.Println("this bullet is not exist.")
        os.Exit(1)
    }

    contents, err := ioutil.ReadFile(fp)
    if err != nil {
        fmt.Println("Can't read bullet file. ", err)
        os.Exit(1)
    }
    fmt.Println(string(contents))
}

