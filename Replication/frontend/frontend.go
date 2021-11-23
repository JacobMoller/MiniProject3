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
	defer Conn1.Close()

	FrontendConn2, Conn2 = Dial(8082, name)
	defer Conn2.Close()

	FrontendConn3, Conn3 = Dial(8083, name)
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

	var one, two, three = ValidateResultsReponsesFromServers(responseFromServerOne, responseFromServerTwo, responseFromServerThree)

	var timeleft int64
	if one.TimeLeft > 0 && one.TimeLeft != 30 {
		timeleft = one.TimeLeft
	}
	if two.TimeLeft > 0 && two.TimeLeft != 30 {
		timeleft = two.TimeLeft
	}
	if three.TimeLeft > 0 && three.TimeLeft != 30 {
		timeleft = three.TimeLeft
	}

	var bidder string
	if one.Bidder != "" {
		bidder = one.Bidder
	}
	if two.Bidder != "" {
		bidder = two.Bidder
	}
	if three.Bidder != "" {
		bidder = three.Bidder
	}

	var serverHighestBid = MaxInt(one.Amount, two.Amount, three.Amount)
	if timeleft > 0 {
		if in.Amount > serverHighestBid {
			FrontendConn1.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: in.Bidder, Amount: in.Amount})
			FrontendConn2.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: in.Bidder, Amount: in.Amount})
			FrontendConn3.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: in.Bidder, Amount: in.Amount})
			return &protobuf.NewBidReply{Message: "Your bid was confirmed."}, nil
		} else {
			return &protobuf.NewBidReply{Message: "Your bid is lower or the same as the current bid. Current bid is " + strconv.FormatInt(serverHighestBid, 10)}, nil //TODO: errors.New("Your bid is lower or the same as the current bid.")
		}
	}
	return &protobuf.NewBidReply{Message: "Time is up. Winner is \"" + bidder + "\" with " + strconv.FormatInt(serverHighestBid, 10)}, nil //TODO: errors.New("Time is up")
}

func ValidateResultsReponsesFromServers(responseFromServerOne, responseFromServerTwo, responseFromServerThree *protobuf.ResultReply) (*protobuf.ResultReply, *protobuf.ResultReply, *protobuf.ResultReply) {
	var one, two, three *protobuf.ResultReply
	if responseFromServerOne.String() != "<nil>" {
		one = responseFromServerOne
	}
	if responseFromServerTwo.String() != "<nil>" {
		two = responseFromServerTwo
	}
	if responseFromServerThree.String() != "<nil>" {
		three = responseFromServerThree
	}
	return one, two, three
}

func ValidateTimeReponsesFromServers(responseFromServerOne, responseFromServerTwo, responseFromServerThree *protobuf.GetTimeReply) (int64, int64, int64) {
	var one, two, three int64 = 30, 30, 30
	if responseFromServerOne.String() != "<nil>" {
		one = responseFromServerOne.TimeLeft
	}
	if responseFromServerTwo.String() != "<nil>" {
		two = responseFromServerTwo.TimeLeft
	}
	if responseFromServerThree.String() != "<nil>" {
		three = responseFromServerThree.TimeLeft
	}
	return one, two, three
}

func MaxInt(x, y, z int64) int64 {
	if x >= y && x >= z {
		return x
	} else if y >= x && y >= z {
		return y
	}
	return z
}

func MinInt(x, y, z int64) int64 {
	var min int64
	if x < y && x < z && x > 0 {
		min = x
	} else if y < x && y < z && y > 0 {
		min = y
	} else if z > 0 {
		min = z
	}
	return min
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
	var responseFromServerOne, _ = FrontendConn1.Result(context.Background(), &protobuf.ResultRequest{})   //200
	var responseFromServerTwo, _ = FrontendConn2.Result(context.Background(), &protobuf.ResultRequest{})   //200
	var responseFromServerThree, _ = FrontendConn3.Result(context.Background(), &protobuf.ResultRequest{}) //150

	var one, two, three = ValidateResultsReponsesFromServers(responseFromServerOne, responseFromServerTwo, responseFromServerThree)
	var bidder, amount, timeleft = BringResultToSync(one, two, three)

	return &protobuf.ResultReply{Bidder: bidder, Amount: amount, TimeLeft: timeleft}, nil
}

func BringResultToSync(one, two, three *protobuf.ResultReply) (string, int64, int64) {
	var bidder string
	var serverHighestBid = MaxInt(one.Amount, two.Amount, three.Amount)
	var timeleft int64
	if one.Amount != two.Amount || two.Amount != three.Amount || one.Amount != three.Amount {
		//Override all values to bring to sync
		FrontendConn1.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: one.Bidder, Amount: serverHighestBid})
		FrontendConn2.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: two.Bidder, Amount: serverHighestBid})
		FrontendConn3.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: three.Bidder, Amount: serverHighestBid})
	}
	if one.Bidder != "" {
		bidder = one.Bidder
	}
	if two.Bidder != "" {
		bidder = two.Bidder
	}
	if three.Bidder != "" {
		bidder = three.Bidder
	}

	if one.TimeLeft > 0 && one.TimeLeft != 30 {
		timeleft = one.TimeLeft
	}
	if two.TimeLeft > 0 && two.TimeLeft != 30 {
		timeleft = two.TimeLeft
	}
	if three.TimeLeft > 0 && three.TimeLeft != 30 {
		timeleft = three.TimeLeft
	}

	return bidder, serverHighestBid, timeleft
}

func (s *server) GetTime(ctx context.Context, in *protobuf.GetTimeRequest) (*protobuf.GetTimeReply, error) {
	var responseFromServerOne, _ = FrontendConn1.GetTime(context.Background(), &protobuf.GetTimeRequest{})
	var responseFromServerTwo, _ = FrontendConn2.GetTime(context.Background(), &protobuf.GetTimeRequest{})
	var responseFromServerThree, _ = FrontendConn3.GetTime(context.Background(), &protobuf.GetTimeRequest{})

	var one, two, three = ValidateTimeReponsesFromServers(responseFromServerOne, responseFromServerTwo, responseFromServerThree)
	var time = BringTimeToSync(one, two, three)

	return &protobuf.GetTimeReply{TimeLeft: time}, nil
}

func BringTimeToSync(one, two, three int64) int64 {
	var minimumTimeLeft = MinInt(one, two, three)
	fmt.Println("Minimum: " + strconv.FormatInt(minimumTimeLeft, 10) + " Based on " + strconv.FormatInt(one, 10) + "; " + strconv.FormatInt(two, 10) + "; " + strconv.FormatInt(three, 10) + ";")
	if one != two || two != three || one != three {
		//Override all values to bring to sync
		FrontendConn1.NewTime(context.Background(), &protobuf.NewTimeRequest{TimeLeft: minimumTimeLeft})
		FrontendConn2.NewTime(context.Background(), &protobuf.NewTimeRequest{TimeLeft: minimumTimeLeft})
		FrontendConn3.NewTime(context.Background(), &protobuf.NewTimeRequest{TimeLeft: minimumTimeLeft})
	}
	return minimumTimeLeft
}
