package models

type Post struct {
  PostId         int
  ThreadId       int
  UserId         int
  Body           string
  CreatedAt      int64

  // Transient
  Thread         *Thread
  User           *User
}
