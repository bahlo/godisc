package models

import (
  "fmt"
  "github.com/coopernurse/gorp"
)

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

func (b Post) String() string {
  return fmt.Sprintf("Post(%s)", b.User)
}

func (b *Post) PreInsert(_ gorp.SqlExecutor) error {
  b.ThreadId = b.Thread.ThreadId
  b.UserId = b.User.UserId

  return nil
}

func (b *Post) PostGet(exe gorp.SqlExecutor) error {
  var (
    obj interface{}
    err error
  )

  obj, err = exe.Get(User{}, b.UserId)
  if err != nil {
    return fmt.Errorf("failed loading a posts's user (%d): %s", b.UserId, err)
  }
  b.User = obj.(*User)

  obj, err = exe.Get(Thread{}, b.ThreadId)
  if err != nil {
    return fmt.Errorf("failed loading a posts's thread (%d): %s", b.ThreadId, err)
  }
  b.Thread = obj.(*Thread)

  return nil
}
