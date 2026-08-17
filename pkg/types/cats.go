package types

import "time"

type Cat struct {
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Port      int    `json:"port"` // gRpc port

	TotalCpu float64 `json:"total_cpu"`
	TotalRam int64   `json:"total_ram"`
	RamUsage float64 `json:"ram_usage"`
	CpuUsage float64 `json:"cpu_Usage"`

	Status        string    `json:"status"` //Active, Asleep, Dead
	LastHeartbeat time.Time `json:"last_heartbeat"`
}
