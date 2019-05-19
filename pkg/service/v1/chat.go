package v1


type chatService struct {
	msg chan string
}

func NewChatServiceServer() v1.ChatServiceServer {
	return &chatService{msg: make(chan string, 1000)}
}

func ()  {
	
}
