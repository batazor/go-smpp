package grpc

import (
	"context"
	"github.com/batazor/go-smpp/pb"
)

func (s *server) SendMessage(ctx context.Context, in *command.MessageSendRequest) (*command.MessageResponse, error) {
	//var response mongo.Document

	//mongo.AddMessageCh <- mongo.Document{
	//	Item: in.Msg,
	//	Response: response,
	//}

	return nil, nil
}

func (s *server) StatusMessage(ctx context.Context, in *command.MessageIdRequest) (*command.MessageResponse, error) {
	var Messages command.MessageResponse

	//mongo.StatusMessageCh <- mongo.Document{
	//	ID: &in.MessageId,
	//	Callback: func(i interface{}, arr []interface{}, err error) {
	//		for _,k := range arr {
	//			Messages.Messages = append(Messages.Messages, &command.Message{
	//				Payload: k.([]byte),
	//			})
	//		}
	//	},
	//}

	return &Messages, nil
}

func (s *server) HistoryMessage(ctx context.Context, in *command.MessageIdRequest) (*command.MessageResponse, error) {
	//mongo.Hist

	return nil, nil
}
