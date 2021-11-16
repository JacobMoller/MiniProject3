package FrontEnd

import (
	"MiniProject3/Replication/protobuf"
	"context"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

type FrontEnd struct {
	Name string
	Port int
}

func New(name string, port int) *FrontEnd {
	fmt.Println("New FE!")

	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	frontend := protobuf.NewReplicationClient(conn)
	message, err2 := frontend.NewNode(context.Background(), &protobuf.NewNodeRequest{Name: name, Type: *protobuf.NewNodeRequest_FrontEnd.Enum()})
	if err2 != nil {
		//Error handling
		if message == nil {
			fmt.Println("Username is already in use for this type")
			return nil
		}
	} else {
		//Start to do stuff here
	}
	return &FrontEnd{name, port}
}

func (f *FrontEnd) ToString() {
	fmt.Println(f.Name + " " + strconv.Itoa(f.Port))
}
