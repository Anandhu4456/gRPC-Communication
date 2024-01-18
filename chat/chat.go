package chat

import (
	"context"
	"log"
)

type Server struct{}

// mustEmbedUnimplementedChatSeviceServer implements ChatSeviceServer.
func (*Server) mustEmbedUnimplementedChatSeviceServer() {
	panic("unimplemented")
}

func (s *Server) SayHello(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("Recieved message body from client %s", msg.Body)
	return &Message{
		Body: "Hello from the Server ",
	}, nil
}
