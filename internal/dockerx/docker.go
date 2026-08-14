package dockerx

//renamed because go compling

import (
	"context"

	"github.com/moby/moby/api/types/container"
)

type Engine interface {
	Pull(ctx context.Context, image string) (string, error)
	Create(ctx context.Context, image string, hostConfig container.HostConfig) (string, error)
	Start(ctx context.Context, container string) (string, error)
	Stop(ctx context.Context, container string) (string, error) // may need to add a timeout for chill shutdown
	Kill(ctx context.Context, container string) (string, error) // after no heartbeat for 300 seconds(5 minutes kill that boy and delte em)
}

type fakeDocker struct{}

func (f fakeDocker) Pull(ctx context.Context, image string) (string, error) {
	return "nginx:alpine", nil
}

func (f fakeDocker) Create(ctx context.Context, image string, hostConfig container.HostConfig) (string, error) {
	return "paw-123 Created!", nil
}

func (f fakeDocker) Start(ctx context.Context, image string) (string, error) {
	return "starting paw-123...", nil
}

func (f fakeDocker) Stop(ctx context.Context, image string) (string, error) {
	return "stopping paw-123...", nil
}

func (f fakeDocker) Kill(ctx context.Context, image string) (string, error) {
	return "killing paw-123... x_x", nil
}
