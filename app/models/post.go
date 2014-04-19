package models

import (
  "time"
)

type Post struct {
  PostId         int
  ThreadId       int
  UserId         int
  Body           string
  CreatedAt      time.Time

  // Transient
  Thread         *Thread
  User           *User
}
