package controllers

import (
  "github.com/revel/revel"
)

type Threads struct {
  *revel.Controller
  GorpController
}

// Index
func (c Threads) Index() revel.Result {
  return c.Render()
}
