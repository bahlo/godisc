package controllers

import (
  "github.com/revel/revel"
  "godisc/app/routes"
)

type App struct {
  *revel.Controller
}

func (c App) Index() revel.Result {
  // TODO: Redirect to index if logged in
  return c.Redirect(routes.App.ShowLogin())
}

func (c App) ShowLogin() revel.Result {
  return c.Render()
}
