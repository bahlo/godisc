package models

import (
  "fmt"
  "time"
  "github.com/coopernurse/gorp"
)

type Post struct {
  PostId         int
  ThreadId       int
  UserId         int
  Body           string
  CreatedString  string

  // Transient
  Thread         *Thread
  User           *User
  Created        interface{}
}

func (b Post) String() string {
  return fmt.Sprintf("Post(%s)", b.User)
}

func (b *Post) PreInsert(_ gorp.SqlExecutor) error {
  b.ThreadId = b.Thread.ThreadId
  b.UserId = b.User.UserId

  b.CreatedString = b.Created.(time.Time).Format(SQL_DATE_FORMAT)

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

  if b.Created, err = time.Parse(SQL_DATE_FORMAT, b.CreatedString); err != nil {
    return fmt.Errorf("failed parsing check in date '%s':", b.CreatedString, err)
  }

  return nil
}
