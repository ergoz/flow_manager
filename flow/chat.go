package flow

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/webitel/flow_manager/model"
)

func (r *router) broadcastChatMessage(ctx context.Context, scope *Flow, conn model.Connection, args interface{}) (model.Response, *model.AppError) {
	var err *model.AppError
	var argv = model.BroadcastChat{}
	var typeProfile string

	if err = scope.Decode(args, &argv); err != nil {
		return nil, err
	}

	if len(argv.Peer) == 0 {
		return nil, ErrorRequiredParameter("broadcastChatMessage", "peer")
	}

	if argv.Profile.Id > 0 {
		typeProfile, err = r.fm.ChatProfileType(conn.DomainId(), argv.Profile.Id)
	}

	peer := make([]model.BroadcastPeer, 0, len(argv.Peer))
	for _, v := range argv.Peer {
		switch p := v.(type) {
		case string:
			peer = append(peer, model.BroadcastPeer{
				Id:   p,
				Type: typeProfile,
				Via:  fmt.Sprintf("%d", argv.Profile.Id),
			})
		case map[string]any:
			peer = append(peer, model.BroadcastPeer{
				Id:   scope.parseString(model.StringValueFromMap("id", p, "")),
				Type: scope.parseString(model.StringValueFromMap("type", p, "")),
				Via:  scope.parseString(model.StringValueFromMap("via", p, "")),
			})
		}
	}

	resp, err := r.fm.BroadcastChatMessage(ctx, conn.DomainId(), argv, peer)
	if err != nil {
		return nil, err
	}
	if len(resp.Failed) != 0 && (argv.FailedReceivers != "" || argv.ResponseCode != "") {
		// save previous logic with response code saved from first peer error message
		status, err := conn.Set(ctx, model.Variables{
			argv.ResponseCode: resp.Failed[0].Error,
		})
		if err != nil {
			return status, err
		}

		// new logic when all failed receivers saved to the variable
		bytes, commonError := json.Marshal(resp)
		if commonError != nil {
			return nil, model.NewAppError("", "flow.chat.broadcast_chat_message.marshal_failed.marshal_error", nil, commonError.Error(), http.StatusInternalServerError)
		}
		status, err = conn.Set(ctx, model.Variables{
			argv.FailedReceivers: string(bytes),
		})
		if err != nil {
			return status, err
		}
	}

	// if the chat_manager service wants to set new variables let him do this
	if len(resp.Variables) != 0 {
		vars := make(model.Variables)
		for key, value := range resp.Variables {
			vars[key] = value
		}
		return conn.Set(ctx, vars)
	}

	return model.CallResponseOK, nil
}
