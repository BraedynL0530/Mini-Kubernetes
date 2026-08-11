package meowlet

import (
	"context"
	"sync"

	"github.com/BraedynL0530/Mini-Kubernetes/internal/dockerx"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/types"
)

type Server struct {
	docker dockerx.Engine
	paw    map[string]*types.Paw
	mu     sync.Mutex
}

func (s *Server) StartPaw(ctx *context.Context, req *pb.StartPawRequest) (resp *pb.StartPawResponse, err error) {
	container, err := s.docker.Create(ctx, req.Image, config)
	if err != nil {
		return nil, err
	}
	s.docker.Start(ctx, container)
	return resp, nil
}
