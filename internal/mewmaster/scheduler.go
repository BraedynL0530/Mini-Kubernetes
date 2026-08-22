package mewmaster

import (
	"context"

	"github.com/BraedynL0530/Mini-Kubernetes/internal/config"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
	"github.com/redis/go-redis/v9"
)

//sechudal will decied what needs to be done,
// helper funcs are self explaintory
// will use concurency/worker pools for each task/ execution.

var ctx = context.Background() //prolly need this
type scheduler struct {
	grpcClient   map[string]pb.MeowletServiceClient
	desiredState *config.ClusterConfig //desired state
	currentState *redis.Client
}

func (s *scheduler) Start() { //these two will be worker/helper funcs
}

func (s *scheduler) Stop() {}

func (s *scheduler) GetCurrenState() map[string]string {
	// use scan to get all keys and values
	// then return that
	// prolly gonna be a map of structs if possible /later
}
func schedule() {

}
