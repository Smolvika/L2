package main

type Handler1 struct {
	Next Handler
}

func (h *Handler1) SendRequest(msg int) string {
	if msg == 1 {
		return "Handler №1"
	} else if h.Next != nil {
		return h.Next.SendRequest(msg)
	}
	return ""
}
