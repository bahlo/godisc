package models

import (
  "github.com/revel/revel"
  "time"
)

type Post struct {
  PostId         int
  ThreadId       int
  UserId         int
  Body           string

  // Transient
  Post           *Post
  Thread         *Thread
  User           *User
}
