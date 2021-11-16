package main

import (
	"MiniProject3/Replication/protobuf"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"bufio"
	"strings"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedReplicationServer
}

type Server struct {
	Name string
	Port int
}

var frontends []string
var servers []string
var primary string

func main() {
	log.Print("Welcome Server. You need to provide a name (either S1, S2 or S3):")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.Replace(text, "\n", "", 1)
	port := strings.Replace(name, "S", "", 1)
	primary = "S1"

	lis, err := net.Listen("tcp", ":808"+port)

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	protobuf.RegisterReplicationServer(s, &server{})

	if (name == primary) {
		//Open port or something stuff out
		go testPrimary()
	} else{
		//Be ready to listen to the primary
		go notPrimary()
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { //error while listening
		log.Fatalf("failed to serve: %v", err)
	}
}

func testPrimary(){
	fmt.Println("Primary. Setting up dial")

	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	frontend := protobuf.NewReplicationClient(conn)
	message, err2 := frontend.NewNode(context.Background(), &protobuf.NewNodeRequest{Name: "Secret-Server Talking :)", Type: *protobuf.NewNodeRequest_Server.Enum()})
	if err2 != nil {
		//Error handling
		if message == nil {
			fmt.Println("Username is already in use for this type")
		}
	} else {
		//Start to do stuff here
		//client.something()
		fmt.Println("Dial Done")
	}
}

func notPrimary(){
	fmt.Println("Not primary. Setting up listener")
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
	}
	printSlice(frontends)
	printSlice(servers)
	//Register the new FrontEnd
	//Broadcast this new info to all Servers
	//After reply from servers, reply to FrontEnd
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

func (s *Server) ToString() {
	fmt.Println(s.Name + " " + strconv.Itoa(s.Port))
}