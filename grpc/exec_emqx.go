package grpc

import (
	"context"
	"log"
	"net/http"
	"strings"

	gm "github.com/log-agent/emqx-grpc/grpc/proto"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/sys"
)

func (s *ManagerServer) ExecEmqxControl(ctx context.Context, in *gm.EmqxRequest) (*gm.EmqxResponse, error) {
	ok := &gm.EmqxResponse{Status: http.StatusOK, Message: "ok", Data: nil}
	log.Println("grpc name:", in.SerName)
	log.Println("grpc active:", in.Active)
	switch in.Active {
	case "start":
		err := sys.CmdRun("docker", in.Active, in.SerName)
		if err != nil {
			ok.Status = http.StatusBadRequest
			ok.Message = err.Error()
		}
	case "stop":
		err := sys.CmdRun("docker", in.Active, in.SerName)
		if err != nil {
			ok.Status = http.StatusBadRequest
			ok.Message = err.Error()
		}
	case "status":
		// docker ps | grep container-name
		output, err := sys.CmdOutTrim("docker", "ps")
		// log.Println(output)
		// log.Println(err)
		if err != nil {
			ok.Status = http.StatusBadRequest
			ok.Message = err.Error()
		} else if !strings.Contains(output, in.SerName) {
			ok.Status = http.StatusBadRequest
			ok.Message = in.SerName + " is stoped"
		}
	default:
		ok.Message = str.Join("操作类型错误:", in.Active)
	}
	return ok, nil
}
