package models

type User struct {
  UserId         int
  Name           string
  HashedPassword []byte
}
