package controllers

import (
  "fmt"
  "time"
  "sort"
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

// Set up sorting for threads
type SortThreads []*models.Thread
func (s SortThreads) Len() int           { return len(s) }
func (s SortThreads) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortThreads) Less(i, j int) bool {
  return s[i].Created.(time.Time).Unix() > s[j].Created.(time.Time).Unix()
}

// Set up sorting for Posts
type SortPosts []*models.Post
func (s SortPosts) Len() int           { return len(s) }
func (s SortPosts) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortPosts) Less(i, j int) bool {
  return s[i].Created.(time.Time).Unix() > s[j].Created.(time.Time).Unix()
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

  sort.Sort(SortThreads(threads))
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
    0,
    topic,
    "",
    user,
    time.Now(),
  }
  err := c.Txn.Insert(thread)
  if err != nil {
    c.Flash.Error("An error occurred, sorry")
    fmt.Println(err)
  }

  return c.Redirect("/threads/%d", thread.ThreadId)
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

  sort.Sort(SortPosts(posts))

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
        "",
        thread,
        user,
        time.Now(),
      }

      err := c.Txn.Insert(post)
      if err != nil {
        c.Flash.Error("An error occurred, sorry")
        fmt.Println(err)
      }
    }

    return c.Redirect("/threads/%d", thread.ThreadId)
  }

  return c.Todo()
}

func (c Threads) DeletePost(id, postId int) revel.Result {
  // Find post
  post, _ := c.Txn.Get(models.Post{}, postId)
  user := c.connected()

  if user.UserId == post.(*models.Post).User.UserId {
    _, err := c.Txn.Delete(post)

    if err != nil {
      c.Flash.Error("An error occurred, sorry")
      fmt.Println(err)
    }
  }

  return c.Redirect("/threads/%d", id)
}
