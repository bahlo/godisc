package models

import (
  "time"
)

type Thread struct {
  ThreadId       int
  UserId         int
  Topic          string
  CreatedAt      time.Time

  // Transient
  User           *User
}
