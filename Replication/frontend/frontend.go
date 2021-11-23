package main

import (
	"MiniProject3/Replication/protobuf"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Welcome Frontend. You need to provide a name for the server to remember you:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.Replace(text, "\n", "", 1)

	var FrontendConn1, Conn1 = Dial(8081, name)
	go SendBids(FrontendConn1)
	defer Conn1.Close()

	var FrontendConn2, Conn2 = Dial(8082, name)
	go SendBids(FrontendConn2)
	defer Conn2.Close()

	var FrontendConn3, Conn3 = Dial(8083, name)
	go SendBids(FrontendConn3)
	defer Conn3.Close()

	//Listen for client bids here

	time.Sleep(1000 * time.Second)
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
