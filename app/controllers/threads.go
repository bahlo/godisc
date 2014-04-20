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
  results, err := c.Txn.Select(models.Thread{},
    `SELECT * FROM Thread`)
  if err != nil {
    // TODO: Catch it
    panic(err)
  }

  var threads []*models.Thread
  for _, r := range results {
    t := r.(*models.Thread)
    threads = append(threads, t)
  }

  return c.Render(threads)
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
    0, // TODO: Fix this: c.connected().UserId,
    topic,
    time.Now().Unix(),
    nil,
  }
  err := c.Txn.Insert(thread)
  if err != nil {
    c.Flash.Error("An error occurred, sorry")
    fmt.Println(err)
  }

  return c.Redirect(Threads.ShowNew)
}

func (c Threads) getThread(id int) *models.Thread {
  threads, err := c.Txn.Select(models.Thread{},
    `SELECT * FROM Thread WHERE ThreadId = ?`, id)
  // TODO: Get posts

  if err != nil {
    // TODO: User not found message
    panic(err)
  }

  // No user found
  if len(threads) == 0 {
    return nil
  }

  return threads[0].(*models.Thread)
}

func (c Threads) Show(id int) revel.Result {
  thread := c.getThread(id)

  return c.Render(thread)
}
