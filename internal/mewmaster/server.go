package mewmaster

import (
	"context"
	"time"

	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/types"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Rbd *redis.Client
}

func (s *Server) RegisterCat(ctx context.Context, req *pb.RegisterCatsRequest) (resp *pb.RegisterCatResponse, err error) {

	err = s.Rbd.HSet(ctx, req.Name, map[string]interface{}{
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

	catData := types.Cat{
		Name:     req.Name,
		CpuUsage: req.CpuUsage,
		RamUsage: req.RamUsage, //error will be fixxed when proto are updated
		Status:   "Active",
		//LastHeartbeat: time.Now().Unix(), //later
	}
	err = s.Rbd.HSet(ctx, req.Name, catData{
		"cpuUsage": req.CpuUsage,
		"ramUsage": req.RamUsage,
		"status":   "Active", // need a time since last heartbeat in other
	}).Err() // err fix this with package or helper func still needs time since last heartbeat

	if err != nil {
		return nil, err
	}

	err = s.Rbd.Expire(ctx, req.Name, 300*time.Second).Err()

	if err != nil {
		//dunno if ill do that or not beccause it needs an error but maybe log instead of return !
		return nil, err
	}

	resp = &pb.CatHeartbeatResponse{Acknowledged: true}
	return resp, nil
}
