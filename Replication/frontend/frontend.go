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

const amountOfServers int = 3

var FrontEndConns [amountOfServers]protobuf.ReplicationClient
var Conns [amountOfServers]*grpc.ClientConn

func main() {
	log.Print("Welcome Frontend. You need to provide a name for the server to remember you:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.Replace(text, "\n", "", 1)

	go FrontendServerStart()
	for i := 0; i < amountOfServers; i++ {
		portToDial := 8080 + i + 1
		FrontEndConns[i], Conns[i] = Dial(portToDial, name)
		defer Conns[i].Close()
	}

	time.Sleep(1000 * time.Second)
}

func FrontendServerStart() {
	lis, err := net.Listen("tcp", ":8085")

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	protobuf.RegisterReplicationServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Dial(port int, name string) (protobuf.ReplicationClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(":"+strconv.Itoa(port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	frontend := protobuf.NewReplicationClient(conn)
	message, userAlreadyExistsError := frontend.NewNode(context.Background(), &protobuf.NewNodeRequest{Name: name, Type: *protobuf.NewNodeRequest_FrontEnd.Enum()})
	if userAlreadyExistsError != nil {
		if message == nil {
			fmt.Println("Username is already in use")
		}
	} else {
		fmt.Println("Dial to " + strconv.Itoa(port) + " was succesful")
		return frontend, conn
	}
	return nil, nil
}

func GetBids(frontend protobuf.ReplicationClient) {
	frontend.Result(context.Background(), &protobuf.ResultRequest{})
}

func (s *server) NewNode(ctx context.Context, in *protobuf.NewNodeRequest) (*protobuf.NewNodeReply, error) {
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
	responsesFromServers := make([]*protobuf.ResultReply, amountOfServers)
	for i := 0; i < len(FrontEndConns); i++ {
		responsesFromServers[i], _ = FrontEndConns[i].Result(context.Background(), &protobuf.ResultRequest{})
	}

	var validatedResultsFromServers = ValidateResultsReponsesFromServers(responsesFromServers)

	var timeleft int64
	var bidder string
	amounts := make([]int64, amountOfServers)
	for i := 0; i < len(validatedResultsFromServers); i++ {
		var current = validatedResultsFromServers[i]
		if current != nil {
			if current.TimeLeft > 0 && current.TimeLeft != 30 {
				timeleft = current.TimeLeft
			}
			if current.Bidder != "" {
				bidder = current.Bidder
			}
			amounts[i] = current.Amount
		}
	}

	var serverHighestBid = MaxInt(amounts)
	if timeleft > 0 {
		if in.Amount > serverHighestBid {
			for i := 0; i < len(FrontEndConns); i++ {
				FrontEndConns[i].NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: in.Bidder, Amount: in.Amount})
			}
			return &protobuf.NewBidReply{Message: "Your bid was confirmed."}, nil
		} else {
			return &protobuf.NewBidReply{Message: "Your bid is lower or the same as the current bid. Current bid is " + strconv.FormatInt(serverHighestBid, 10)}, nil
		}
	}
	return &protobuf.NewBidReply{Message: "Time is up. Winner is \"" + bidder + "\" with " + strconv.FormatInt(serverHighestBid, 10)}, nil
}

func ValidateResultsReponsesFromServers(responsesFromServers []*protobuf.ResultReply) []*protobuf.ResultReply {
	validatedResultsFromServers := make([]*protobuf.ResultReply, amountOfServers)
	for i := 0; i < len(responsesFromServers); i++ {
		if responsesFromServers[i].String() != "<nil>" {
			validatedResultsFromServers[i] = responsesFromServers[i]
		}
	}
	return validatedResultsFromServers
}

func ValidateTimeReponsesFromServers(responsesFromServers []*protobuf.GetTimeReply) []int64 {
	timesLeft := []int64{30, 30, 30}

	for i := 0; i < len(responsesFromServers); i++ {
		if responsesFromServers[i].String() != "<nil>" {
			timesLeft[i] = responsesFromServers[i].TimeLeft
		}
	}
	return timesLeft
}

func MaxInt(amounts []int64) int64 {
	var highestAmount int64
	for i := 0; i < len(amounts); i++ {
		if amounts[i] > highestAmount {
			highestAmount = amounts[i]
		}
	}
	return highestAmount
}

func MinInt(amounts []int64) int64 {
	var lowestAmount int64 = 31
	for i := 0; i < len(amounts); i++ {
		if lowestAmount > amounts[i] {
			lowestAmount = amounts[i]
		}
	}
	return lowestAmount
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

func (s *server) Result(ctx context.Context, in *protobuf.ResultRequest) (*protobuf.ResultReply, error) {
	responsesFromServers := make([]*protobuf.ResultReply, amountOfServers)

	for i := 0; i < len(FrontEndConns); i++ {
		responsesFromServers[i], _ = FrontEndConns[i].Result(context.Background(), &protobuf.ResultRequest{})
	}

	var validatedResultsFromServers = ValidateResultsReponsesFromServers(responsesFromServers)
	var bidder, amount, timeleft = BringResultToSync(validatedResultsFromServers)

	return &protobuf.ResultReply{Bidder: bidder, Amount: amount, TimeLeft: timeleft}, nil
}

func BringResultToSync(validatedResultsFromServers []*protobuf.ResultReply) (string, int64, int64) {
	var bidder string
	amounts := []int64{-1, -1, -1}
	for i := 0; i < len(validatedResultsFromServers); i++ {
		var current = validatedResultsFromServers[i]
		if current != nil {
			amounts[0] = current.Amount
		}
	}

	var serverHighestBid = MaxInt(amounts)
	var nameOfHighestBidder string
	for i := 0; i < len(validatedResultsFromServers); i++ {
		if validatedResultsFromServers[i] != nil && validatedResultsFromServers[i].Amount == serverHighestBid {
			nameOfHighestBidder = validatedResultsFromServers[i].Bidder
		}
	}
	var timeleft int64
	if !IsEqual(amounts) {
		for i := 0; i < len(amounts); i++ {
			if amounts[i] != -1 {
				FrontEndConns[i].NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: nameOfHighestBidder, Amount: serverHighestBid})
			}
		}
	}
	for i := 0; i < len(validatedResultsFromServers); i++ {
		var current = validatedResultsFromServers[i]
		if current != nil {
			if current.Bidder != "" {
				bidder = current.Bidder
			}
			if current.TimeLeft > 0 && current.TimeLeft != 30 {
				timeleft = current.TimeLeft
			}
		}
	}

	return bidder, serverHighestBid, timeleft
}

func IsEqual(amounts []int64) bool {
	var firstAmount = amounts[0]
	for i := 1; i < len(amounts); i++ {
		if firstAmount != amounts[i] {
			return false
		}
	}
	return true
}

func (s *server) GetTime(ctx context.Context, in *protobuf.GetTimeRequest) (*protobuf.GetTimeReply, error) {
	responsesFromServers := make([]*protobuf.GetTimeReply, amountOfServers)
	for i := 0; i < amountOfServers; i++ {
		responsesFromServers[i], _ = FrontEndConns[i].GetTime(context.Background(), &protobuf.GetTimeRequest{})
	}

	var validatedResultsFromServers = ValidateTimeReponsesFromServers(responsesFromServers)
	var time = BringTimeToSync(validatedResultsFromServers)

	return &protobuf.GetTimeReply{TimeLeft: time}, nil
}

func BringTimeToSync(validatedResultsFromServers []int64) int64 {
	var minimumTimeLeft = MinInt(validatedResultsFromServers)
	fmt.Println("Minimum: " + strconv.FormatInt(minimumTimeLeft, 10) + " Based on " + strconv.FormatInt(validatedResultsFromServers[0], 10) + "; " + strconv.FormatInt(validatedResultsFromServers[1], 10) + "; " + strconv.FormatInt(validatedResultsFromServers[2], 10) + ";")
	if !IsEqual(validatedResultsFromServers) {
		for i := 0; i < len(FrontEndConns); i++ {
			FrontEndConns[i].NewTime(context.Background(), &protobuf.NewTimeRequest{TimeLeft: minimumTimeLeft})
		}
	}
	return minimumTimeLeft
}
