package model

// Response is the response to the api request
// for node information.
type InfoResponse struct {
	ID              int     `db:"id" json:"-"`
	Blocks          int     `db:"blocks" json:"blocks"`
	Connections     int     `db:"connections" json:"connections"`
	Country         string  `db:"country" json:"country"`
	Difficulty      float64 `db:"difficulty" json:"difficulty"`
	DonationAddress string  `db:"donationAddress" json:"donationAddress"`
	IP              string  `db:"ip" json:"ip"`
	IncomingData    float64 `db:"incomingData" json:"incomingData"`
	MaxFee          float64 `db:"maxFee" json:"maxFee"`
	MaxMemory       float64 `db:"maxMemory" json:"maxMemory"`
	MidFee          float64 `db:"midFee" json:"midFee"`
	MinFee          float64 `db:"minFee" json:"minFee"`
	Network         string  `db:"network" json:"network"`
	NetworkHashPS   float64 `db:"networkHashPS" json:"networkHashPS"`
	OutgoingData    float64 `db:"outgoingData" json:"outgoingData"`
	Protocol        int     `db:"protocol" json:"protocol"`
	Rank            int     `db:"rank" json:"rank"`
	Status          string  `db:"status" json:"status"`
	StakingStatus   string  `db:"stakingStatus" json:"stakingStatus"`
	Subversion      string  `db:"subversion" json:"subversion"`
	TotalData       float64 `db:"totalData" json:"totalData"`
	Transactions    int     `db:"transactions" json:"transactions"`
	UsedMemory      float64 `db:"usedMemory" json:"usedMemory"`
	Version         int     `db:"version" json:"version"`
}
