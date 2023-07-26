package main

type Handler3 struct {
	Next Handler
}

func (h *Handler3) SendRequest(msg int) string {
	if msg == 3 {
		return "Handler №3"
	} else if h.Next != nil {
		return h.Next.SendRequest(msg)
	}
	return "Data entered incorrectly"
}
