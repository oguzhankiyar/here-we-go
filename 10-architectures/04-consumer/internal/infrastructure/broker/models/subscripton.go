package models

type Subscription struct {
	Id string
	Fn func([]byte)
}
