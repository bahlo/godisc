package controllers

import (
  "github.com/revel/revel"
)

type Settings struct {
  App
  *revel.Controller
}

func (s Settings) Index() revel.Result {
  return s.Render()
}

func (s Settings) Save(username string) revel.Result {
  return s.Redirect(s.Index)
}
