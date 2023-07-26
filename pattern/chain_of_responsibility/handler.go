package main

type Handler interface {
	SendRequest(msg int) string
}
