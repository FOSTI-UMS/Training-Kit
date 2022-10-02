package main

import (
  "log"
  "github.com/sethvargo/go-password/password"
)

func generate() {
  res, err := password.Generate(64, 10, 10, false, false)
  if err != nil {
    log.Fatal(err)
  }
  log.Printf(res)
}