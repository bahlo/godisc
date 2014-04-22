package controllers

import (
  "time"
  "github.com/revel/revel"
)

func init() {
  revel.OnAppStart(InitDB)
  revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
  revel.InterceptMethod(App.AddUser, revel.BEFORE)
  revel.InterceptMethod(App.AddConfig, revel.BEFORE)
  revel.InterceptMethod(Threads.checkUser, revel.BEFORE)
  revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
  revel.InterceptMethod((*GorpController).Rollback, revel.PANIC)

  revel.TemplateFuncs["eqo"] = func(a, b, c interface{}) bool {
    return a == b || a == c
  }
  revel.TemplateFuncs["formatDate"] = func(t time.Time) string {
    if format, ok := revel.Config.String("date.format"); ok {
      return t.Format(format)
    }

    return "2006-01-02 15:04:05"
  }
}
