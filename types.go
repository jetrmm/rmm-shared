package jetrmm

type CheckInNats struct {
	AgentId string `json:"agent_id"`
	Version string `json:"version"`
}

type AgentInfoNats struct {
	AgentId      string  `json:"agent_id"`
	Username     string  `json:"logged_in_username"`
	Hostname     string  `json:"hostname"`
	OS           string  `json:"operating_system"`
	Platform     string  `json:"plat"`
	TotalRAM     float64 `json:"total_ram"`
	BootTime     int64   `json:"boot_time"`
	RebootNeeded bool    `json:"needs_reboot"`
	GoArch       string  `json:"goarch"`
}

type Storage struct {
	Device  string `json:"device"`
	Fstype  string `json:"fstype"`
	Total   string `json:"total"`
	Used    string `json:"used"`
	Free    string `json:"free"`
	Percent int    `json:"percent"`
}

type PublicIPNats struct {
	AgentId  string `json:"agent_id"`
	PublicIP string `json:"public_ip"`
}
