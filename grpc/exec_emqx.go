package grpc

import (
	"context"
	"net/http"

	gm "github.com/log-agent/emqx-grpc/grpc/proto"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/sys"
)

func (s *ManagerServer) ExecEmqxControl(ctx context.Context, in *gm.EmqxRequest) (*gm.EmqxResponse, error) {
	ok := &gm.EmqxResponse{Status: http.StatusOK, Message: "ok", Data: nil}
	switch in.Active {
	case "start":
	case "stop":
		if err := sys.CmdRun("docker", in.Active, in.SerName); err != nil {
			ok.Status = http.StatusBadRequest
			ok.Message = err.Error()
		}
	default:
		ok.Message = str.Join("操作类型错误:", in.Active)
	}
	return ok, nil
}
