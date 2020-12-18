package model

type Message struct {
	Body        string
	Subject     string
	ContentType string
	To          []string
}
