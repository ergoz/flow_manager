package app

import (
	"context"
	"github.com/webitel/call_center/grpc_api/cc"
	"github.com/webitel/flow_manager/model"
)

func (fm *FlowManager) JoinToInboundQueue(ctx context.Context, in *cc.CallJoinToQueueRequest) (cc.MemberService_CallJoinToQueueClient, error) {
	return fm.cc.Member().JoinCallToQueue(ctx, in)
}

func (fm *FlowManager) AddMemberToQueueQueue(domainId int64, queueId int, number, name string, typeId, holdSec int, variables map[string]string) *model.AppError {
	return fm.Store.Call().AddMemberToQueueQueue(domainId, queueId, number, name, typeId, holdSec, variables)
}
