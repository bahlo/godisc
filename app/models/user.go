package models

import (
  "github.com/revel/revel"
)

type User struct {
  UserId         int
  Name           string
  HashedPassword []byte
}
