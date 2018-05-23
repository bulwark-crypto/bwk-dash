package rpc

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// RPC is a wrapper around the Coin struct
// that is the connection to the node.
type RPC struct {
	node *Coin
}

// Call will make a call to the rpc node.
func (client *RPC) Call(method string, args ...interface{}) (data []byte, err error) {
	data, err = client.node.Call(method, args...)
	return
}

// GetFees will return the max, mid, and min fees.
func (client *RPC) GetFees() (max, mid, min float64) {
	var data []byte
	var err error
	if data, err = client.node.Call("estimatefee", 1); err != nil {
		fmt.Println(err.Error())
		return
	}
	fees := new(GetFees)
	if err = json.Unmarshal(data, &fees); err != nil {
		fmt.Println(err.Error())
		return
	}
	min = fees.Result

	if data, err = client.node.Call("estimatefee", 12); err != nil {
		fmt.Println(err.Error())
		return
	}
	fees = new(GetFees)
	if err = json.Unmarshal(data, &fees); err != nil {
		fmt.Println(err.Error())
		return
	}
	mid = fees.Result

	if data, err = client.node.Call("estimatefee", 25); err != nil {
		fmt.Println(err.Error())
		return
	}
	fees = new(GetFees)
	if err = json.Unmarshal(data, &fees); err != nil {
		fmt.Println(err.Error())
		return
	}
	max = fees.Result
	return
}

// GetInfo will return the getinfo data as a struct.
func (client *RPC) GetInfo() (info *GetInfo, err error) {
	var data []byte
	data, err = client.node.Call("getinfo")
	if err != nil {
		return
	}

	info = new(GetInfo)
	err = json.Unmarshal(data, info)
	return
}

// GetNetTotals will return the bytes used on the network
// for the node.
func (client *RPC) GetNetTotals() (recv, sent float64) {
	recv = 0
	sent = 0

	var data []byte
	var err error
	if data, err = client.node.Call("getnettotals"); err != nil {
		fmt.Println(err.Error())
		return
	}

	res := new(GetNetTotals)
	if err = json.Unmarshal(data, res); err != nil {
		fmt.Println(err.Error())
		return
	}

	recv = res.Result.Recv
	sent = res.Result.Sent
	return
}

// GetNetworkHashPS will return the current network hash.
func (client *RPC) GetNetworkHashPS() (hashps float64) {
	hashps = 0 // Default to 0

	var data []byte
	var err error
	data, err = client.node.Call("getnetworkhashps")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := new(GetNetworkHashPS)
	if err = json.Unmarshal(data, res); err != nil {
		fmt.Println(err.Error())
		return
	}

	hashps = res.Result
	return
}

// GetTransactions will return the total number of transactions.
func (client *RPC) GetTransactions() (count int) {
	count = 0

	var data []byte
	var err error
	if data, err = client.node.Call("gettxoutsetinfo"); err != nil {
		fmt.Println(err.Error())
		return
	}

	res := new(GetTransactions)
	if err = json.Unmarshal(data, res); err != nil {
		fmt.Println(err.Error())
		return
	}

	count = res.Result.Transactions
	return
}

// GetUsedMemory will return the bytes of used memory.
func (client *RPC) GetUsedMemory() (size float64) {
	size = 0

	var data []byte
	var err error
	if data, err = client.node.Call("getmempoolinfo"); err != nil {
		fmt.Println(err.Error())
		return
	}

	res := new(GetUsedMemory)
	if err = json.Unmarshal(data, res); err != nil {
		fmt.Println(err.Error())
		return
	}

	size = res.Result.Bytes
	return
}

// GetVersion will return the subversion from getnetworkinfo.
func (client *RPC) GetVersion() (version string) {
	var data []byte
	var err error
	if data, err = client.node.Call("getnetworkinfo"); err != nil {
		fmt.Println(err.Error())
		return
	}

	res := new(GetNetworkInfo)
	if err = json.Unmarshal(data, res); err != nil {
		fmt.Println(err.Error())
		return
	}

	version = res.Result.Subversion
	return
}

// NewRPC returns a new rpc client with a connection
// to a node included internally.
func NewRPC() (rpc *RPC, err error) {
	// For now rpc connection information is
	// loaded from ENV.
	rpcaddr := os.Getenv("DASH_RPC_ADDR")
	rpcuser := os.Getenv("DASH_RPC_USER")
	rpcpass := os.Getenv("DASH_RPC_PASS")
	rpcport, err := strconv.ParseInt(os.Getenv("DASH_RPC_PORT"), 10, 64)
	if err != nil {
		fmt.Println("Warning: unable to parse port from .env, setting to default.")
		rpcport = 52541
	}

	// Setup the rpc client.
	rpc = new(RPC)
	rpc.node, err = NewCoin(rpcuser, rpcpass, rpcaddr, "", int(rpcport))
	if err != nil {
		return
	}

	return
}
