package rpc

// Credit goes to github.com/HarleyLiu/btcRPC
// Thanks for the work in providing the base.

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// CoinDefaultHost   default coin address
const CoinDefaultHost string = "localhost"

//CoinDefaultPort   default coin port
const CoinDefaultPort int = 52541 // Bulwark

//CoinDefaultProto  default coin protp
const CoinDefaultProto string = "http"

//RPCTimeOut default timeout(second)
const RPCTimeOut = 10

//Coin RPC struct
type Coin struct {
	// Configuration options
	username    string
	password    string
	proto       string
	host        string
	port        int
	url         string
	certificate string
	// Information and debugging
	Status    int
	LastError error
	// rawResponse string
	responseData []byte
	id           int
	client       *http.Client
}

//NewCoin   create a new RPC instance
func NewCoin(coinUser, coinPasswd, coinHost, coinURL string, coinPort int) (cn *Coin, err error) {
	cn = &Coin{
		username: coinUser,
		password: coinPasswd,
		host:     coinHost,
		port:     coinPort,
		url:      coinURL,
		proto:    CoinDefaultProto,
	}
	if len(coinHost) == 0 {
		cn.host = CoinDefaultHost
	}
	if coinPort < 0 || coinPort > 65535 {
		cn.port = CoinDefaultPort
	}
	cn.client = &http.Client{}
	cn.client.Timeout = time.Duration(RPCTimeOut) * time.Second
	cn.client.Transport = &http.Transport{}
	//first access
	if _, err = cn.Call("getinfo"); err != nil {
		return nil, err
	}
	if cn.Status != http.StatusOK || cn.LastError != nil {
		return nil, cn.LastError
	}
	return cn, nil
}

//SetSSL    setup certificate
func (cn *Coin) SetSSL(certificate string) {
	cn.proto = "https"
	cn.certificate = certificate
}

func (cn *Coin) access(data map[string]interface{}) (err error) {
	if len(data) != 2 {
		err = errors.New("params count error")
		return
	}
	if cn.client == nil {
		err = errors.New("http client error")
		return
	}
	cn.id++
	data["id"] = cn.id
	cn.LastError = nil
	cn.responseData = nil
	cn.Status = http.StatusOK
	var (
		jbuf []byte
		req  *http.Request
		resp *http.Response
	)
	if jbuf, err = json.Marshal(data); err != nil {
		return
	}

	addr := cn.proto + "://" + cn.username + ":" + cn.password + "@" + cn.host + ":" + strconv.Itoa(cn.port) + "/" + cn.url
	if req, err = http.NewRequest("POST", addr, bytes.NewReader(jbuf)); err != nil {
		cn.LastError = err
		return
	}
	req.Header.Set("Content-Type", "application/json")
	//todo: setup ssl
	if resp, err = cn.client.Do(req); err != nil {
		cn.LastError = err
		return
	}
	cn.Status = resp.StatusCode
	if cn.Status != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	defer resp.Body.Close()
	if cn.responseData, err = ioutil.ReadAll(resp.Body); err != nil {
		cn.LastError = err
		return
	}
	if len(cn.responseData) == 0 {
		err = errors.New("response data is empty")
		return
	}
	return
}

//Call run RPC command
func (cn *Coin) Call(method string, args ...interface{}) (data []byte, err error) {
	if method == "" {
		err = errors.New("method is not set")
		return
	}
	requestData := make(map[string]interface{})
	requestData["method"] = method
	requestData["params"] = args
	if err = cn.access(requestData); err == nil {
		data = cn.responseData
	}
	return
}
