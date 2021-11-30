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

var commandList = "\n To bid, enter an integer.\n To see the current highest bid, enter \"get\"\n To see the time left, enter \"time\""
var name string

func main() {
	log.Print("Welcome Client. You need to provide a name for the frontend to remember you:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name = strings.Replace(text, "\n", "", 1)

	conn, err := grpc.Dial(":8085", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := protobuf.NewReplicationClient(conn)
	message, err2 := client.NewNode(context.Background(), &protobuf.NewNodeRequest{Name: name, Type: *protobuf.NewNodeRequest_Client.Enum()})
	if err2 != nil {
		//Error handling
		if message == nil {
			fmt.Println("Username is already in use")
		}
	} else {
		fmt.Println("Welcome to the auction." + commandList)
		go TakeInput(client)
		time.Sleep(1000 * time.Second)
	}
}

func TakeInput(client protobuf.ReplicationClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputParsed := strings.Replace(input, "\n", "", 1)
		amount, err := strconv.ParseInt(inputParsed, 10, 64)
		if err != nil {
			if inputParsed == "get" {
				var result, _ = client.Result(context.Background(), &protobuf.ResultRequest{})
				fmt.Println("Current highest bid is: " + strconv.FormatInt(result.Amount, 10) + " and is from \"" + result.Bidder + "\"")
			} else if inputParsed == "time" {
				var result, _ = client.GetTime(context.Background(), &protobuf.GetTimeRequest{})
				fmt.Println("Current time left: " + strconv.FormatInt(result.TimeLeft, 10))
				//method here
			} else {
				fmt.Println("Unknown command" + commandList)
			}
		} else {
			var result, _ = client.NewBid(context.Background(), &protobuf.NewBidRequest{Bidder: name, Amount: amount})
			fmt.Println(result.Message)
		}
	}
}
