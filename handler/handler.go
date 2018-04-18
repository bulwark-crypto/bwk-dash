package handler

import (
	"github.com/dustinengle/bwk-dash/rpc"
	"github.com/gin-gonic/gin"
)

// GetNodeInfo will return all the information
// needed to be displayed on the dashboard.
func GetNodeInfo(c *gin.Context) {
	// Get node rpc connection from request context.
	node := c.MustGet("rpc").(*rpc.RPC)

	// Get node information.
	info, err := node.GetInfo()
	if err != nil {
		c.Error(err)
		return
	}

	// Setup the response object.
	res := new(InfoResponse)
	res.Blocks = info.Result.Blocks
	res.Blocks = info.Result.Blocks
	res.Connections = info.Result.Connections
	res.Country = ""
	res.Difficulty = info.Result.Difficulty
	res.DonationAddress = c.MustGet("donation").(string)
	res.IP = c.MustGet("ip").(string)
	res.MaxMemory = 0 // TODO: get system memory for mempool bar.
	res.NetworkHashPS = node.GetNetworkHashPS()
	res.Protocol = info.Result.Protocol
	res.Rank = 0 // TODO: get masternode rank if masternode.
	res.Status = "Up"
	res.StakingStatus = info.Result.StakingStatus
	res.Subversion = node.GetVersion()
	res.Transactions = node.GetTransactions()
	res.UsedMemory = 0
	res.Version = info.Result.Version

	// Setup the network name.
	res.Network = "mainnet"
	if info.Result.Testnet {
		res.Network = "testnet"
	}

	// Get the network traffic.
	res.IncomingData, res.OutgoingData = node.GetNetTotals()
	res.TotalData = res.IncomingData + res.OutgoingData

	// Get the max, mid, and min fees.
	res.MaxFee, res.MidFee, res.MinFee = node.GetFees()

	// Get memory information for the node.
	res.UsedMemory = node.GetUsedMemory()

	// Return as json to client.
	c.JSON(200, res)
}
