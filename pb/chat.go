package pb

import (
	"log"

	"golang.org/x/net/context"
)

//Server this is simple type
type Server struct {
}

//SayHello this is first basic funtion
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Println("Received body from the client: ", message.Body)
	return &Message{Body: "Message from the server!"}, nil
}
