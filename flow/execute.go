package flow

import "github.com/webitel/flow_manager/model"

type executeArgs struct {
	flow *Flow
}

func (r *router) execute(c model.Connection, args interface{}) (model.Response, *model.AppError) {
	return ResponseOK, nil
}
