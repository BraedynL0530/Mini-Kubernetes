package types

import "time"

type Cat struct {
	Name string `json:"name"`
	IpAdress string `json:"ip_adress"`
	Port int `json:"port"` // gRpc port
	Status string `json:"status"` //Active, Alseep, Dead
	LastHeartbeat time.Time `json:"last_heartbeat"`
}
