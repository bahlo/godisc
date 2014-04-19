package controllers

import (
  "fmt"
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/go-sql-driver/mysql"
  "github.com/revel/revel"
  "github.com/revel/revel/modules/db/app"
  "godisc/app/models"
)

var (
  Dbm *gorp.DbMap
)

func InitDB() {
  db.Init()
  Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

  // Users
  Dbm.AddTable(models.User{}).SetKeys(true, "UserId")

  // Posts
  t := Dbm.AddTable(models.Post{}).SetKeys(true, "PostId")
  t.ColMap("User").Transient = true
  t.ColMap("Thread").Transient = true

  // Threads
  t = Dbm.AddTable(models.Thread{}).SetKeys(true, "ThreadId")
  t.ColMap("User").Transient = true

  Dbm.TraceOn("[gorp]", revel.INFO)
  Dbm.CreateTablesIfNotExists()

  fmt.Println("finished InitDB()")
}

type GorpController struct {
  *revel.Controller
  Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }
  c.Txn = txn
  return nil
}

func (c *GorpController) Commit() revel.Result {
  if c.Txn == nil {
    return nil
  }
  if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Txn = nil
  return nil
}

func (c *GorpController) Rollback() revel.Result {
  if c.Txn == nil {
    return nil
  }
  if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Txn = nil
  return nil
}
