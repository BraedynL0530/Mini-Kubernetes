package mewmaster

import (
	"context"

	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
)

type Server struct {
}

func (s *Server) RegisterCat(ctx context.Context, req *pb.RegisterCatsRequest) (resp *pb.RegisterCatResponse, err error) {
	//need a cache, lets use redis

	return resp, nil
}
