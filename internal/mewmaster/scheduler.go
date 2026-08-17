package mewmaster

import (
	"github.com/BraedynL0530/Mini-Kubernetes/internal/config"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/proto/pb"
)

conn, err := grpc.NewClient(nodeAddr, grpc.WithInsecure())// look up uh, credentials and stuff!
client := pb.MeowletServiceClient(conn)
defer conn.Close()
//got connection, now lets schedule
// gotta get the stuff from config,
// then dynamically do it,
// add math for catsurge based on resource usage in nodes, paw usage/load
// and yeah, desired state and stuff
// lets do concurency! this is going to need multiple connections for each daemon.
type sechudlar struct {
	client map[string]pb.MeowletServiceClient
	config *config.ClusterConfig


}
func schedule() {

}
