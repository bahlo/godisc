package controllers

import (
  "fmt"
  "time"
  "github.com/revel/revel"
  "godisc/app/routes"
  "godisc/app/models"
)

type Threads struct {
  App
  *revel.Controller
  GorpController
}

func (c Threads) checkUser() revel.Result {
  if user := c.connected(); user == nil {
    c.Flash.Error("Please log in first")

    return c.Redirect(routes.App.Index())
  }

  return nil
}

// Index shows all threads
func (c Threads) Index() revel.Result {
  return c.Render()
}

// ShowNew shows the form for new threads
func (c Threads) ShowNew() revel.Result {
  return c.Render()
}

// New creates a new thread
func (c Threads) New(topic string) revel.Result {
  // Validate
  c.Validation.Required(topic)
  c.Validation.MinSize(topic, 3)

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(Threads.ShowNew)
  }

  // Create thread
  thread := &models.Thread{
    0,
    0,
    topic,
    time.Now(),
    nil,
  }
  err := c.Txn.Insert(thread)
  if err != nil {
    c.Flash.Error("An error occurred, sorry")
    fmt.Println(err)
  }

  return c.Redirect(Threads.ShowNew)
}
