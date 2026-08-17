package meowlet

import (
	"context"
	"sync"

	"github.com/BraedynL0530/Mini-Kubernetes/internal/dockerx"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/types"
	"github.com/moby/moby/api/types/container"
)

//TODO: figure out actual id! req.pawid LIKELY WONT WORK!!!!!

type Server struct {
	docker dockerx.Engine // do i even need pull image? well see later choom.
	paw    map[string]*types.Paw
	mu     sync.Mutex
}

func (s *Server) StartPaw(ctx context.Context, req *pb.StartPawRequest) (resp *pb.StartPawResponse, err error) {
	config := container.HostConfig{}
	paw, err := s.docker.Create(ctx, req.Image, config) // paw is container i just cannot use container keyword already!

	if err != nil {
		return nil, err
	}
	s.docker.Start(ctx, paw)

	resp = &pb.StartPawResponse{Success: true}
	return resp, nil
}

func (s *Server) StopPaw(ctx context.Context, req *pb.StopPawRequest) (resp *pb.StopPawResponse, err error) {
	_, err = s.docker.Stop(ctx, req.PawId)

	if err != nil {
		return nil, err
	}

	_, err = s.docker.Kill(ctx, req.PawId) //Note: id  is seperate from container id unless i change it to match
	if err != nil {
		return nil, err
	}

	resp = &pb.StopPawResponse{Success: true}
	return resp, nil
}

// why am i talking to myself in third person?
