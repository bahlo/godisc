package controllers

import (
  "github.com/revel/revel"
  // "godisc/app/models"
)

type Settings struct {
  App
  *revel.Controller
}

func (c Settings) Index() revel.Result {
  if user := c.connected(); user == nil {
    return c.Redirect(App.ShowLogin)
  }
  return c.Render()
}
