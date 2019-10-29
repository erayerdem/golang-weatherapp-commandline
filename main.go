package main

import (
  "fmt"
  "github.com/urfave/cli"
  "os"
  "weather/commands"
)

func main() {
  app := cli.NewApp()
  app.Name = "Weather App"
  app.Usage = "Learn Weather basic way"
  app.Version = "1.0.0"
  app.HelpName = "it is basic"
  app.Commands = commands.GetCommands()

  err := app.Run(os.Args)
  if err != nil {
    fmt.Print(err)
    os.Exit(1)
  }
}
