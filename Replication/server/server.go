package main

import (
	"MiniProject3/Replication/protobuf"
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedReplicationServer
}

var frontends []string

func main() {
	lis, err := net.Listen("tcp", ":8080")

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

func (s *server) NewFrontEnd(ctx context.Context, in *protobuf.NewFrontEndRequest) (*protobuf.NewFrontEndReply, error) {
	//Register the new FrontEnd

	if alreadyExists(in.FrontEndName) {
		fmt.Println("FrontEnd DENIED (frontEndName: \"" + in.FrontEndName + "\")")
		return &protobuf.NewFrontEndReply{}, errors.New("USERNAME IS ALREADY IN USE")
	} else {
		fmt.Println("NEW FrontEnd (frontEndName: \"" + in.FrontEndName + "\")")
		frontends = append(frontends, in.FrontEndName)
	}
	//Broadcast this new info to all Servers
	//After reply from servers, reply to FrontEnd
	return &protobuf.NewFrontEndReply{}, nil
}

func alreadyExists(frontEndName string) bool {
	var existsInFrontEnd = false
	for i := 0; i < len(frontends); i++ {
		if frontends[i] == frontEndName {
			existsInFrontEnd = true
		}
	}
	return existsInFrontEnd
}
