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
  user := c.connected()
  thread := &models.Thread{
    0,
    user.UserId,
    topic,
    time.Now().Unix(),
    user,
  }
  err := c.Txn.Insert(thread)
  if err != nil {
    c.Flash.Error("An error occurred, sorry")
    fmt.Println(err)
  }

  return c.Redirect(Threads.Index)
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

func (c Threads) getPosts(threadId int) []*models.Post {
  posts, err := c.Txn.Select(models.Post{},
    `SELECT * FROM Post WHERE ThreadId = ?`, threadId)

  if err != nil {
    panic(err)
  }

  if len(posts) == 0 {
    return nil
  }

  tc := []*models.Post{}
  for _, post := range posts {
    tc = append(tc, post.(*models.Post))
  }

  return tc
}

func (c Threads) Show(id int) revel.Result {
  thread := c.getThread(id)
  posts := c.getPosts(id)

  return c.Render(thread, posts)
}

func (c Threads) Post(id int, body string) revel.Result {
  if len(body) > 0 {
    thread := c.getThread(id)
    user := c.connected()
    if thread != nil && user != nil {
      post := &models.Post{
        0,
        thread.ThreadId,
        user.UserId,
        body,
        time.Now().Unix(),
        thread,
        user,
      }

      err := c.Txn.Insert(post)
      if err != nil {
        c.Flash.Error("An error occurred, sorry")
        fmt.Println(err)
      }
    }
  }

  return c.Redirect(Threads.Index)
}
