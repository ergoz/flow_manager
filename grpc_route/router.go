package grpc_route

import (
	"context"
	"fmt"
	"github.com/webitel/flow_manager/app"
	"github.com/webitel/flow_manager/flow"
	"github.com/webitel/flow_manager/model"
	"net/http"
)

type Router struct {
	fm   *app.FlowManager
	apps flow.ApplicationHandlers
}

func Init(fm *app.FlowManager, fr flow.Router) {
	var router = &Router{
		fm: fm,
	}

	router.apps = flow.UnionApplicationMap(
		ApplicationsHandlers(router),
		fr.Handlers(),
	)

	fm.GRPCRouter = router
}

func (r *Router) Request(ctx context.Context, scope *flow.Flow, req model.ApplicationRequest) <-chan model.Result {
	if h, ok := r.apps[req.Id()]; ok {
		if h.ArgsParser != nil {
			return h.Handler(ctx, scope, h.ArgsParser(scope.Connection, req.Args()))

		} else {
			return h.Handler(ctx, scope, req.Args())
		}
	} else {
		return flow.Do(func(result *model.Result) {
			result.Err = model.NewAppError("GRPC.Request", "grpc.request.not_found", nil, fmt.Sprintf("appId=%v not found", req.Id()), http.StatusNotFound)
		})
	}
}

func (r *Router) Handle(conn model.Connection) *model.AppError {

	gr := conn.(model.GRPCConnection)

	s, err := r.fm.GetSchemaById(conn.DomainId(), gr.SchemaId())
	if err != nil {
		return err
	}

	i := flow.New(flow.Config{
		Name:     s.Name,
		Schema:   s.Schema,
		Handler:  r,
		Conn:     conn,
		Timezone: "",
	})

	flow.Route(conn.Context(), i, r)

	return nil
}

func (r *Router) Decode(scope *flow.Flow, in interface{}, out interface{}) *model.AppError {
	return scope.Decode(in, out)
}
