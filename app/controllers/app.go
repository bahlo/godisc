package controllers

import (
  "code.google.com/p/go.crypto/bcrypt"
  "github.com/revel/revel"
  "godisc/app/models"
)

type App struct {
  *revel.Controller
  GorpController
}

// Index
func (c App) Index() revel.Result {
  if user := c.connected(); user != nil {
    return c.Redirect("/threads")
  }
  return c.Redirect(App.ShowLogin)
}

// Show login form
func (c App) ShowLogin() revel.Result {
  if user := c.connected(); user != nil {
    return c.Redirect("/threads")
  }
  return c.Render()
}

func (c App) AddUser() revel.Result {
  if user := c.connected(); user != nil {
    c.RenderArgs["user"] = user
  }

  return nil
}

func (c App) AddConfig() revel.Result {
  if name, ok := revel.Config.String("app.name"); ok {
    c.RenderArgs["name"] = name
  }

  return nil
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

func (c App) Login(username, password string) revel.Result {
  user := c.getUser(username)

  if user != nil {
    err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))

    if err == nil {
      c.Session["user"] = username
      c.Session.SetDefaultExpiration()

      return c.Redirect(Threads.Index)
    }
  }

  // Set flash cookie
  c.Flash.Out["username"] = username
  c.Flash.Error("Login failed, wrong Username or Password combination")

  return c.Redirect(App.ShowLogin)
}

func (c App) Logout() revel.Result {
  for k := range c.Session {
    delete(c.Session, k)
  }

  return c.Redirect(App.Index)
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
