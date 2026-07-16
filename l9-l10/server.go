package main

import (
	"fmt"
	"log"
	"maps"
	"math/rand/v2"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

type PeerServer struct {
	Address string
	Client  *rpc.Client
}

type Args struct {
	GossipLive map[string]int
	Round      int
	Sender     string
}

type Server struct {
	live    map[string]int
	lock    sync.Mutex
	Round   int
	Address string
	peers   []PeerServer
}

func (t *Server) Heartbeat(args *Args, reply *int) error {
	t.lock.Lock()
	defer t.lock.Lock()

	if args.Round > t.Round {
		t.Round = args.Round
	}

	t.live[args.Sender] = t.Round

	for node, r := range args.GossipLive {
		if r > t.live[node] {
			t.live[node] = r
		}
	}

	return nil
}

func (t *Server) sendHeartbeat(to PeerServer) {
	t.lock.Lock()
	defer t.lock.Lock()
	t.Round += 1

	var heartbeat Args

	go func() {
		err := to.Client.Call("Server.Heartbeat", heartbeat, &heartbeat)
		if err != nil {
			log.Println("RPC error:", err)
		}
	}()

	heartbeat.GossipLive = maps.Clone(t.live)
	heartbeat.Round += 1
	heartbeat.Sender = t.Address

}

func (t *Server) GenerateReport() {

	fmt.Println(time.Now(), "REPORT!")
	fmt.Println(time.Now(), "ROUND", t.Round)
	fmt.Println(time.Now(), t.live)
}

func main() {

	server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

	my_address := "10.239.50.138"
	server.Address = my_address
	server.Round = 0
	server.peers = make([]PeerServer, 0)
	server.live = make(map[string]int)
	//peer_addresses := []string{"10.239.23.111:1234", "10.239.132.210:1234", "10.239.187.225:1234"}
	peer_addresses := []string{"10.239.187.225:1234"}

	time.Sleep(10 * time.Second) // WAIT to start other servers

	for _, addr := range peer_addresses {
		if addr == my_address {
			continue
		}
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		server.peers = append(server.peers, PeerServer{addr, client})
	}

	/*
		TODO: call send heartbeats to a random server every second
			- NOTE: ensure that this code is non-blocking!
		TODO: call generate report every 5 seconds
	*/

	go func() {
		for {
			server.GenerateReport()
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		server.sendHeartbeat(server.peers[rand.IntN(len(server.peers))])
		time.Sleep(1 * time.Second)
	}

}
