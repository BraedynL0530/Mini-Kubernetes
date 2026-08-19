package mewmaster

import (
	"context"
	"time"

	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/types"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	rbd *redis.Client
}

func (s *Server) RegisterCat(ctx context.Context, req *pb.RegisterCatsRequest) (resp *pb.RegisterCatResponse, err error) {

	err = s.rbd.HSet(ctx, req.Name, map[string]interface{}{
		"ip":       req.Ip,
		"totalCpu": req.TotalCpu,
		"totalRam": req.TotalRam,
		"status":   "Active", // need a time since last heartbeat in other
	}).Err()

	if err != nil {
		return nil, err
	}

	resp = &pb.RegisterCatResponse{Success: true}
	return resp, nil
}

func (s *Server) CatHeartbeat(ctx context.Context, req *pb.CatHeartbeatRequest) (resp *pb.CatHeartbeatResponse, err error) {

	err = s.rbd.HSet(ctx, req.Name, structs.Map[string](types.Cat)).Err() // err fix this with package or helper func still needs time since last heartbeat
	if err != nil {
		return nil, err
	}
	err = s.rbd.Expire(ctx, req.Name, 300*time.Second).Err()
	if err != nil {
		//dunno if ill do that or not beccause it needs an error but maybe log instead of return !
		return nil, err
	}

	resp = &pb.CatHeartbeatResponse{Acknowledged: true}
	return resp, nil
}
