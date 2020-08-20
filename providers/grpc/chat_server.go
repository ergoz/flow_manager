package grpc

import (
	"context"
	"fmt"
	"github.com/webitel/engine/utils"
	"github.com/webitel/flow_manager/model"
	"github.com/webitel/flow_manager/providers/grpc/flow"
	"net/http"
)

const (
	activeConversationCacheSize = 50000
	maximumInactiveChat         = 60 * 60 * 24 // day
	confirmationBuffer          = 100
)

type chatApi struct {
	conversations utils.ObjectCache
	*server
}

func NewChatApi(s *server) *chatApi {
	return &chatApi{
		server:        s,
		conversations: utils.NewLru(activeConversationCacheSize),
	}
}

func (s *chatApi) Start(ctx context.Context, req *flow.StartRequest) (*flow.StartResponse, error) {
	if _, ok := s.conversations.Get(req.ConversationId); ok {
		return &flow.StartResponse{
			Error: &flow.Error{
				Id:      "grpc.chat.start.valid.conversation_id",
				Message: fmt.Sprintf("Conversation %d already exists", req.ConversationId),
			},
		}, nil
	}

	conv := NewConversation(req.ConversationId, req.DomainId, req.ProfileId)
	conv.chat = s

	s.conversations.AddWithExpiresInSecs(req.ConversationId, conv, maximumInactiveChat)

	s.server.consume <- conv

	return &flow.StartResponse{}, nil
}

func (s *chatApi) Break(ctx context.Context, req *flow.BreakRequest) (*flow.BreakResponse, error) {
	conv, err := s.getConversation(req.ConversationId)
	if err != nil {
		return &flow.BreakResponse{
			Error: &flow.Error{
				Id:      err.Id,
				Message: err.Message,
			},
		}, nil
	}

	if err := conv.Break(); err != nil {
		return &flow.BreakResponse{
			Error: &flow.Error{
				Id:      err.Id,
				Message: err.Message,
			},
		}, nil
	}

	return &flow.BreakResponse{}, nil
}

func (s *chatApi) ConfirmationMessage(ctx context.Context, req *flow.ConfirmationMessageRequest) (*flow.ConfirmationMessageResponse, error) {
	var conf chan []string
	var ok bool

	conv, err := s.getConversation(req.ConversationId)
	if err != nil {
		return &flow.ConfirmationMessageResponse{
			Error: &flow.Error{
				Id:      err.Id,
				Message: err.Message,
			},
		}, nil
	}

	conv.mx.RLock()
	conf, ok = conv.confirmation[req.ConfirmationId]
	if ok {
		delete(conv.confirmation, req.ConfirmationId)
	}
	conv.mx.Unlock()

	if !ok {
		return &flow.ConfirmationMessageResponse{
			Error: &flow.Error{
				Id:      "chat.grpc.conversation.confirmation.not_found",
				Message: fmt.Sprintf("Confirmation %d not found", req.ConfirmationId),
			},
		}, nil
	}

	msgs := make([]string, len(req.Messages), len(req.Messages))

	// TODO
	for _, m := range req.Messages {
		switch x := m.Value.(type) {
		case *flow.Message_TextMessage_:
			msgs = append(msgs, x.TextMessage.Text)
		}
	}

	conf <- msgs

	return &flow.ConfirmationMessageResponse{}, nil
}

func (s *chatApi) getConversation(id int64) (*conversation, *model.AppError) {
	conv, ok := s.conversations.Get(id)
	if !ok {
		return nil, model.NewAppError("Chat", "grpc.chat.conversation.not_found", nil,
			fmt.Sprintf("Conversation %d not found", id), http.StatusNotFound)
	}

	return conv.(*conversation), nil
}
