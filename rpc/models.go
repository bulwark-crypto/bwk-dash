package rpc

type GetFees struct {
	Error  string  `json:"error"`
	ID     int     `json:"id"`
	Result float64 `json:"result"`
}

type GetInfo struct {
	Error  string `json:"error"`
	ID     int    `json:"id"`
	Result struct {
		Balance            float64 `json:"balance"`
		Blocks             int     `json:"blocks"`
		Connections        int     `json:"connections"`
		Difficulty         float64 `json:"difficulty"`
		Errors             string  `json:"errors"`
		KeyPoolOldest      int     `json:"keypoololdest"`
		KeyPoolSize        int     `json:"keypoolsize"`
		ObfuscationBalance float64 `json:"obfuscation_balance"`
		PayTXFee           float64 `json:"paytxfee"`
		Protocol           int     `json:"protocolversion"`
		Proxy              string  `json:"proxy"`
		RelayFee           float64 `json:"relayfee"`
		StakingStatus      string  `json:"staking status"`
		Testnet            bool    `json:"testnet"`
		TimeOffset         int     `json:"timeoffset"`
		Version            int     `json:"version"`
		Wallet             int     `json:"walletversion"`
	} `json:"result"`
}

type GetNetTotals struct {
	Error  string `json:"error"`
	ID     int    `json:"id"`
	Result struct {
		Recv float64 `json:"totalbytesrecv"`
		Sent float64 `json:"totalbytessent"`
		Time int     `json:"timemillis"`
	} `json:"result"`
}

type GetNetworkHashPS struct {
	Error  string  `json:"error"`
	ID     int     `json:"id"`
	Result float64 `json:"result"`
}

type GetNetworkInfo struct {
	Error  string `json:"error"`
	ID     int    `json:"id"`
	Result struct {
		Version         int           `json:"version"`
		Subversion      string        `json:"subversion"`
		ProtocolVersion int           `json:"protocolversion"`
		LocalServices   string        `json:"localservices"`
		TimeOffset      int           `json:"timeoffset"`
		Connections     int           `json:"connections"`
		Networks        []interface{} `json:"networks"`
		RelayFee        float64       `json:"relayfee"`
		LocalAddresses  []interface{} `json:"localaddresses"`
	} `json:"result"`
}

type GetTransactions struct {
	Error  string `json:"error"`
	ID     int    `json:"id"`
	Result struct {
		Height          int     `json:"height"`
		BestBlock       string  `json:"bestblock"`
		Transactions    int     `json:"transactions"`
		TXOuts          int     `json:"txouts"`
		BytesSerialized float64 `json:"bytes_serialized"`
		HashSerialized  string  `json:"hash_serialized"`
		TotalAmount     float64 `json:"total_amount"`
	} `json:"result"`
}

type GetUsedMemory struct {
	Error  string `json:"error"`
	ID     int    `json:"id"`
	Result struct {
		Bytes float64 `json:"bytes"`
		Size  int     `json:"size"`
	} `json:"result"`
}
