package handler

// Response is the response to the api request
// for node information.
type InfoResponse struct {
	Blocks          int     `json:"blocks"`
	Connections     int     `json:"connections"`
	Country         string  `json:"country"`
	Difficulty      float64 `json:"difficulty"`
	DonationAddress string  `json:"donationAddress"`
	IP              string  `json:"ip"`
	IncomingData    float64 `json:"incomingData"`
	MaxFee          float64 `json:"maxFee"`
	MaxMemory       float64 `json:"maxMemory"`
	MidFee          float64 `json:"midFee"`
	MinFee          float64 `json:"minFee"`
	Network         string  `json:"network"`
	NetworkHashPS   float64 `json:"networkHashPS"`
	OutgoingData    float64 `json:"outgoingData"`
	Protocol        int     `json:"protocol"`
	Rank            int     `json:"rank"`
	Status          string  `json:"status"`
	StakingStatus   string  `json:"stakingStatus"`
	Subversion      string  `json:"subversion"`
	TotalData       float64 `json:"totalData"`
	Transactions    int     `json:"transactions"`
	UsedMemory      float64 `json:"usedMemory"`
	Version         int     `json:"version"`
}
