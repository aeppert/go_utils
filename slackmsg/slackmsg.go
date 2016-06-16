package main

import (
  "os"
  "fmt"
  "github.com/bluele/slack"
)

const (
  token       = "<your token here>"
)

func usage() {
   fmt.Println("usage: slackmsg <username> <channel> <message>");
}

func main() {
  if len(os.Args) < 4 {
    usage();
    os.Exit(1);
  }

  api := slack.New(token)

  channel, err := api.FindChannelByName(os.Args[2])
  if err != nil {
    panic(err)
  }

  err = api.ChatPostMessage(channel.Id, os.Args[3], &slack.ChatPostMessageOpt{ Username: os.Args[1] })
  if err != nil {
    panic(err)
  }
}

