package controllers

import (
  "github.com/revel/revel"
  "godisc/app/routes"
)

type App struct {
  *revel.Controller
  GorpController
}

// Index
func (c App) Index() revel.Result {
  // TODO: Redirect to index if logged in
  // c.Txn.Ping()
  return c.Redirect(routes.App.ShowLogin())
}

// Show login form
func (c App) ShowLogin() revel.Result {
  return c.Render()
}

func (c App) DoLogin(username, password string) revel.Result {
  c.Validation.Required(username)
  c.Validation.Required(password)
  c.Validation.MinSize(username, 3)

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(App.ShowLogin)
  }

  return c.Redirect(routes.App.Index())
}
