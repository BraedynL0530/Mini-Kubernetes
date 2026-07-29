package types

type Paw struct {
	ID string `json:"id"`
	Image string `json:"image"`
	Status string `json:"status"`
	CatIP string `json:"cat_ip"`
	MemRequest int64 `json:"mem_request"`
	CpuRequest float64 `json:"cpu_request"`
	CpuLimit float64 `json:"cpu_limit"`
	MemLimit int64 `json:"mem_limit"`
}
