package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dustinengle/bwk-dash/data"
	"github.com/dustinengle/bwk-dash/handler"
	"github.com/dustinengle/bwk-dash/rpc"
	"github.com/dustinengle/bwk-dash/sys"

	_ "github.com/joho/godotenv/autoload"
)

var (
	db   *data.DB
	node *rpc.RPC
)

// getData will return an info response object
// with the required data to be saved into the
// database for use by API.
func getData() (res *handler.InfoResponse, err error) {
	var info *rpc.GetInfo

	// Get rpc information.
	info, err = node.GetInfo()
	if err != nil {
		return
	}

	// Setup the response object.
	res = new(handler.InfoResponse)
	res.Blocks = info.Result.Blocks
	res.Blocks = info.Result.Blocks
	res.Connections = info.Result.Connections
	res.Country = ""
	res.Difficulty = info.Result.Difficulty
	res.DonationAddress = os.Getenv("DASH_DONATION_ADDRESS")
	res.MaxMemory = sys.GetMemorySize()
	res.NetworkHashPS = node.GetNetworkHashPS()
	res.Protocol = info.Result.Protocol
	res.Rank = 0
	res.Status = "Online"
	res.StakingStatus = info.Result.StakingStatus
	res.Subversion = node.GetVersion()
	res.Transactions = node.GetTransactions()
	res.UsedMemory = 0
	res.Version = info.Result.Version

	// Get the ip address for the node.
	res.IP, err = sys.GetIP()
	if err != nil {
		return
	}

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

	// Get memory information for the rpc.
	res.UsedMemory = node.GetUsedMemory()

	return
}

func main() {
	var err error

	// Setup database connection.
	db, err = data.NewSQL(os.ExpandEnv(os.Getenv("DASH_DB")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup the table structure if not found.
	err = db.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Setup rpc connection to rpc.
	node, err = rpc.NewRPC()
	if err != nil {
		log.Fatal(err)
	}

	// Get information from rpc, apis, etc.
	var info *handler.InfoResponse
	info, err = getData()
	if err != nil {
		log.Fatal(err)
	}

	// Store in database.
	err = db.Save(info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Information saved to database!")
}
