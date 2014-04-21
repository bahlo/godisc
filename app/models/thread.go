package models

import (
  "fmt"
  "github.com/coopernurse/gorp"
)

type Thread struct {
  ThreadId       int
  UserId         int
  Topic          string
  CreatedAt      int64

  // Transient
  User           *User
}

func (b Thread) String() string {
  return fmt.Sprintf("Thread(%s)", b.User)
}

func (b *Thread) PreInsert(_ gorp.SqlExecutor) error {
  b.UserId = b.User.UserId
  return nil
}

func (b *Thread) PostGet(exe gorp.SqlExecutor) error {
  var (
    obj interface{}
    err error
  )

  obj, err = exe.Get(User{}, b.UserId)
  if err != nil {
    return fmt.Errorf("Error loading a threads's user (%d): %s", b.UserId, err)
  }
  b.User = obj.(*User)

  return nil
}
