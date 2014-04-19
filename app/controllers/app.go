package controllers

import (
  "code.google.com/p/go.crypto/bcrypt"
  "github.com/revel/revel"
  "godisc/app/routes"
  "godisc/app/models"
)

type App struct {
  *revel.Controller
  GorpController
}

// Index
func (c App) Index() revel.Result {
  // TODO: Redirect to index if logged in
  return c.Redirect(routes.App.ShowLogin())
}

// Show login form
func (c App) ShowLogin() revel.Result {
  return c.Render()
}

func (c App) getUser(username string) *models.User {
  users, err := c.Txn.Select(models.User{},
    `SELECT * FROM User WHERE Name = ?`, username)

  if err != nil {
    // TODO: User not found message
    panic(err)
  }

  // No user found
  if len(users) == 0 {
    return nil
  }

  return users[0].(*models.User)
}

func (c App) DoLogin(username, password string) revel.Result {
  user := c.getUser(username)

  if user != nil {
    err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))

    if err == nil {
      c.Session["user"] = username
      c.Session.SetDefaultExpiration()

      c.Flash.Success("Welcome, " + username)
      return c.Redirect(routes.Threads.Index())
    }
  }

  c.Flash.Out["username"] = username
  c.Flash.Error("Login failed")
  return c.Redirect(routes.App.Index())
}

func (c App) connected() *models.User {
  if c.RenderArgs["user"] != nil {
    return c.RenderArgs["user"].(*models.User)
  }

  if username, ok := c.Session["user"]; ok {
    return c.getUser(username)
  }

  return nil
}
