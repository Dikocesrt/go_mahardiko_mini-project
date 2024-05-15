package chatbot

type Chatbot struct {
	Role    string
	Content string
}

type UseCaseInterface interface {
	Chat(chat *Chatbot) (*Chatbot, error)
}