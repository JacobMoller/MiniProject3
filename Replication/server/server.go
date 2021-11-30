package main

import (
	"MiniProject3/Replication/protobuf"
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedReplicationServer
}

var frontends []string
var servers []string
var bidder string
var amount int64
var timeleft int64 = -1
var ticking bool

func main() {
	log.Print("Welcome Server. You need to provide a name:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.Replace(text, "\n", "", 1)
	port := strings.Replace(name, "S", "", 1)

	lis, err := net.Listen("tcp", ":808"+port)

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	protobuf.RegisterReplicationServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { //error while listening
		log.Fatalf("failed to serve: %v", err)
	}
}

func TimeTick() {
	ticking = true
	for {
		if timeleft > 0 {
			timeleft--
		} else {
			fmt.Println("Times up")
			break
		}
		fmt.Println("time left: " + strconv.FormatInt(timeleft, 10))
		time.Sleep(time.Second)
	}
}

func (s *server) NewBid(ctx context.Context, in *protobuf.NewBidRequest) (*protobuf.NewBidReply, error) {
	fmt.Println("Server Received bid: " + strconv.FormatInt(in.Amount, 10))
	amount = in.Amount
	bidder = in.Bidder
	return &protobuf.NewBidReply{}, nil
}

func (s *server) NewNode(ctx context.Context, in *protobuf.NewNodeRequest) (*protobuf.NewNodeReply, error) {
	//Which type is this?
	var sliceToUse *[]string
	if in.Type == *protobuf.NewNodeRequest_FrontEnd.Enum() {
		sliceToUse = &frontends
	} else if in.Type == *protobuf.NewNodeRequest_Server.Enum() {
		sliceToUse = &servers
	}
	if alreadyExists(*sliceToUse, in.Name) {
		fmt.Println("Node DENIED (name: \"" + in.Name + "\", type: " + in.Type.String() + ")")
		return &protobuf.NewNodeReply{}, errors.New("USERNAME IS ALREADY IN USE")
	} else {
		fmt.Println("NEW Node (name: \"" + in.Name + "\", type: " + in.Type.String() + ")")
		*sliceToUse = append(*sliceToUse, in.Name)
		if timeleft == -1 {
			timeleft = 30
			go TimeTick()
		}
	}
	printSlice(frontends)
	printSlice(servers)
	//Register the new FrontEnd
	//Broadcast this new info to all Servers
	//After reply from servers, reply to FrontEnd
	return &protobuf.NewNodeReply{}, nil
}

func (s *server) Result(ctx context.Context, in *protobuf.ResultRequest) (*protobuf.ResultReply, error) {
	return &protobuf.ResultReply{Bidder: bidder, Amount: amount, TimeLeft: timeleft}, nil
}

func alreadyExists(pool []string, inputName string) bool {
	var existsInPool = false
	for i := 0; i < len(pool); i++ {
		if pool[i] == inputName {
			existsInPool = true
		}
	}
	return existsInPool
}

func printSlice(sliceToPrint []string) {
	fmt.Print("[")
	for i := 0; i < len(sliceToPrint); i++ {
		fmt.Print(sliceToPrint[i])
		if i != len(sliceToPrint)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
	fmt.Println()
}

func (s *server) GetTime(ctx context.Context, in *protobuf.GetTimeRequest) (*protobuf.GetTimeReply, error) {
	return &protobuf.GetTimeReply{TimeLeft: timeleft}, nil
}

func (s *server) NewTime(ctx context.Context, in *protobuf.NewTimeRequest) (*protobuf.NewTimeReply, error) {
	timeleft = in.TimeLeft
	if !ticking {
		go TimeTick()
	}
	return &protobuf.NewTimeReply{}, nil
}
