package types

type Paw struct {
	ID string `json:"id"`
	Image string `json:"image"`
	Status string `json:"status"`
	CatIP string `json:"cat_ip"`
	CpuCount float64 `json:"cpu_count"`
	MemLimit int64 `json:"mem_limit"`
}
