package models

type Thread struct {
  ThreadId       int
  UserId         int
  Topic          string
  CreatedAt      int64

  // Transient
  User           *User
}
