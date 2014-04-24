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

func (p Post) String() string {
  return fmt.Sprintf("Post(%s)", p.User)
}

func (p *Post) PreInsert(_ gorp.SqlExecutor) error {
  p.ThreadId = p.Thread.ThreadId
  p.UserId = p.User.UserId

  p.CreatedString = p.Created.(time.Time).Format(SQL_DATE_FORMAT)

  return nil
}

func (p *Post) PostGet(exe gorp.SqlExecutor) error {
  var (
    obj interface{}
    err error
  )

  obj, err = exe.Get(User{}, p.UserId)
  if err != nil {
    return fmt.Errorf("failed loading a posts's user (%d): %s", p.UserId, err)
  }
  p.User = obj.(*User)

  obj, err = exe.Get(Thread{}, p.ThreadId)
  if err != nil {
    return fmt.Errorf("failed loading a posts's thread (%d): %s", p.ThreadId, err)
  }
  p.Thread = obj.(*Thread)

  if p.Created, err = time.Parse(SQL_DATE_FORMAT, p.CreatedString); err != nil {
    return fmt.Errorf("failed parsing check in date '%s':", p.CreatedString, err)
  }

  return nil
}
