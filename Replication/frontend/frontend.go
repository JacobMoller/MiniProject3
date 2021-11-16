package main

import (
	"MiniProject3/Replication/protobuf"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Welcome Frontend. You need to provide a name for the server to remember you:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.Replace(text, "\n", "", 1)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	frontend := protobuf.NewReplicationClient(conn)
	message, err2 := frontend.NewFrontEnd(context.Background(), &protobuf.NewFrontEndRequest{FrontEndName: name})
	if err2 != nil {
		//Error handling
		if message == nil {
			fmt.Println("Username is already in use")
		}
	} else {
		//Start to do stuff here
	}
}
