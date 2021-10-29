package main

import (
    "strings"
    "fmt"
    "log"
    "os"
    "github.com/urfave/cli"
)

var app = cli.NewApp()

var pizza = []string{"Enjoy your pizza with some delicious"}

func info() {
    app.Name = "Simple pizza CLI"
    app.Usage = "An example CLI for pizza ordering"
    app.Author = "Son Nguyen"
    app.Version = "1.0.0"
}

func addMaterial(material string) []string {
    return append(pizza, material)
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name: "add:materrial",
            Aliases: []string{"p"},
            Usage: "Add apple to pizza",
            Action: func(c *cli.Context) {
                ap := c.Args().Get(0)
                apples := addMaterial(ap)
                m := strings.Join(apples, " ")
                fmt.Println(m)
            },
        },
    }
}

func main() {
    info()
    commands()

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
