package models

import (
  "github.com/revel/revel"
  "time"
)

type Thread struct {
  ThreadId       int
  UserId         int
  Topic          string

  // Transient
  Thread         *Thread
  User           *User
}
