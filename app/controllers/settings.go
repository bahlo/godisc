package controllers

import (
  "github.com/revel/revel"
)

type Settings struct {
  App
  *revel.Controller
}

func (c Settings) Index() revel.Result {
  return c.Render()
}

func (c Settings) Save(name string) revel.Result {

  if user := c.connected(); len(name) > 0 && user != nil {
    user.Name = name

    c.Session["user"] = name
    c.Session.SetDefaultExpiration()

    c.Txn.Update(user)
  }

  return c.Redirect(Settings.Index)
}
