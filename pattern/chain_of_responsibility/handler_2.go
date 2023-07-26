package main

type Handler2 struct {
	Next Handler
}

func (h *Handler2) SendRequest(msg int) string {
	if msg == 2 {
		return "Handler №2"
	} else if h.Next != nil {
		return h.Next.SendRequest(msg)
	}
	return ""
}
