package main

import (
  "fmt"
  "github.com/urfave/cli/v2" // imports as package "cli"
  "log"
  "os"
  "regexp"
)

func main() {
  app := &cli.App{
    Name: "amazon url shoter",
    Usage: "shoten url",
    Action: func(c *cli.Context) error {
      if c.NArg() > 0 {
        url := c.Args().Get(0)
        ref := regexp.MustCompile(`ref[^]]*`)
        name := regexp.MustCompile(`%[^]]*/dp`)
        url = ref.ReplaceAllString(url, "")
        url = name.ReplaceAllString(url, "dp")
        fmt.Println(url)
      }
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}