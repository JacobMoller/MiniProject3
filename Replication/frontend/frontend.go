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

var clients []string
var FrontendConn1 protobuf.ReplicationClient
var Conn1 *grpc.ClientConn
var FrontendConn2 protobuf.ReplicationClient
var Conn2 *grpc.ClientConn
var FrontendConn3 protobuf.ReplicationClient
var Conn3 *grpc.ClientConn

func main() {
	log.Print("Welcome Frontend. You need to provide a name for the server to remember you:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.Replace(text, "\n", "", 1)

	go FrontendServerStart()

	FrontendConn1, Conn1 = Dial(8081, name)
	//go SendBids(FrontendConn1)
	defer Conn1.Close()

	FrontendConn2, Conn2 = Dial(8082, name)
	//go SendBids(FrontendConn2)
	defer Conn2.Close()

	FrontendConn3, Conn3 = Dial(8083, name)
	//go SendBids(FrontendConn3)
	defer Conn3.Close()

	time.Sleep(1000 * time.Second)
}

func FrontendServerStart() {
	//Listen for client bids here
	//EXPERIMENTAL START
	lis, err := net.Listen("tcp", ":8085")

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	protobuf.RegisterReplicationServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { //error while listening
		log.Fatalf("failed to serve: %v", err)
	}
	//EXPERIMENTAL END
}

func Dial(port int, name string) (protobuf.ReplicationClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(":"+strconv.Itoa(port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	frontend := protobuf.NewReplicationClient(conn)
	message, err2 := frontend.NewNode(context.Background(), &protobuf.NewNodeRequest{Name: name, Type: *protobuf.NewNodeRequest_FrontEnd.Enum()})
	if err2 != nil {
		//Error handling
		if message == nil {
			fmt.Println("Username is already in use")
		}
	} else {
		//Start to do stuff here
		fmt.Println("Dial to " + strconv.Itoa(port) + " was succesful")
		return frontend, conn
	}
	fmt.Println("Returning nil :(")
	return nil, nil
}

func SendBids(frontend protobuf.ReplicationClient) {
	var amount int64 = 0
	for {
		amount++
		fmt.Println("Sending bid: " + strconv.FormatInt(amount, 10))
		frontend.NewBid(context.Background(), &protobuf.NewBidRequest{Amount: amount})
		time.Sleep(2 * time.Second)
	}
}

func GetBids(frontend protobuf.ReplicationClient) {
	frontend.Result(context.Background(), &protobuf.ResultRequest{})
}

func (s *server) NewNode(ctx context.Context, in *protobuf.NewNodeRequest) (*protobuf.NewNodeReply, error) {
	//Which type is this?
	if alreadyExists(clients, in.Name) {
		fmt.Println("Node DENIED (name: \"" + in.Name + "\", type: " + in.Type.String() + ")")
		return &protobuf.NewNodeReply{}, errors.New("USERNAME IS ALREADY IN USE")
	} else {
		fmt.Println("NEW Node (name: \"" + in.Name + "\", type: " + in.Type.String() + ")")
		clients = append(clients, in.Name)
	}
	printSlice(clients)

	return &protobuf.NewNodeReply{}, nil
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

func (s *server) NewBid(ctx context.Context, in *protobuf.NewBidRequest) (*protobuf.NewBidReply, error) {
	fmt.Println("Frontend Received bid: " + strconv.FormatInt(in.Amount, 10))

	//What is the servers current amount????
	var responseFromServerOne, _ = FrontendConn1.Result(context.Background(), &protobuf.ResultRequest{})
	var responseFromServerTwo, _ = FrontendConn2.Result(context.Background(), &protobuf.ResultRequest{})
	var responseFromServerThree, _ = FrontendConn3.Result(context.Background(), &protobuf.ResultRequest{})

	var serverHighestBid = MaxInt(responseFromServerOne.Amount, responseFromServerTwo.Amount, responseFromServerThree.Amount)

	if in.Amount > serverHighestBid {
		fmt.Println("Your bid is very good. tnx 4 moneys")
		FrontendConn1.NewBid(context.Background(), &protobuf.NewBidRequest{Amount: in.Amount})
		FrontendConn2.NewBid(context.Background(), &protobuf.NewBidRequest{Amount: in.Amount})
		FrontendConn3.NewBid(context.Background(), &protobuf.NewBidRequest{Amount: in.Amount})
		return &protobuf.NewBidReply{}, nil
	} else {
		fmt.Println("Your bid is lower or the same as the current bid. Current bid is " + strconv.FormatInt(serverHighestBid, 10))
		return &protobuf.NewBidReply{}, errors.New("Your bid is lower or the same as the current bid. Current bid is " + strconv.FormatInt(serverHighestBid, 10))
	}
}

func MaxInt(x, y, z int64) int64 {
	if x >= y && x >= z {
		return x
	} else if y >= x && y >= z {
		return y
	}
	return z
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
