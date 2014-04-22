package models

import (
  "fmt"
  "time"
  "github.com/coopernurse/gorp"
)

type Thread struct {
  ThreadId        int
  UserId          int
  Topic           string
  CreatedString   string

  // Transient
  User            *User
  Created         interface{}
}

const (
  DATE_FORMAT     = "02. 01. 2006 15:04:05"
  SQL_DATE_FORMAT = "2006-01-02 15:04:05"
)

func (b Thread) String() string {
  return fmt.Sprintf("Thread(%s)", b.User)
}

func (b *Thread) PreInsert(_ gorp.SqlExecutor) error {
  b.UserId = b.User.UserId
  b.CreatedString = b.Created.(time.Time).Format(SQL_DATE_FORMAT)

  return nil
}

func (b *Thread) PostGet(exe gorp.SqlExecutor) error {
  var (
    obj interface{}
    err error
  )

  obj, err = exe.Get(User{}, b.UserId)
  if err != nil {
    return fmt.Errorf("failed loading a threads's user (%d): %s", b.UserId, err)
  }
  b.User = obj.(*User)

  if b.Created, err = time.Parse(SQL_DATE_FORMAT, b.CreatedString); err != nil {
    return fmt.Errorf("failed parsing check in date '%s':", b.CreatedString, err)
  }

  return nil
}
