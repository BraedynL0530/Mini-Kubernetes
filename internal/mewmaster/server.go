package mewmaster

import (
	"context"

	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	rbd *redis.Client
}

func (s *Server) RegisterCat(ctx context.Context, req *pb.RegisterCatsRequest) (resp *pb.RegisterCatResponse, err error) {
	nodeName := req.Name

	err = s.rbd.HSet(ctx, nodeName, map[string]interface{}{
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

	err = s.rbd.HSet(ctx, req.Name, map[string]interface{}{
		"name":     req.Name,
		"cpuUsage": req.CpuUsage,
		"ramUsage": req.RamUsage,
		// add status and time since last heartbeat later
	}).Err()
	if err != nil {
		return nil, err
	}

	resp = &pb.CatHeartbeatResponse{Acknowledged: true}
	return resp, nil
}
