package main

import (
	pb "example/communication"
	"fmt"
	"log"
	"net"

	myerr "example/error"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Define localhost address infomation
const (
	address     = "127.0.0.1"
	defaultname = "server"
	port        = 6666
)

// set ip
var ip = pb.IP{
	Addr: address,
	Name: defaultname,
	Port: port,
}

type server struct{}

// Implement CommunicationServer interface.
func (s *server) Greet(ctxx context.Context, in *pb.GreetRequest) (*pb.GreetReply, error) {
	// Validate GreetRequest
	if err := in.Validate(); err != nil {
		// If validation failed, server return error message.
		log.Printf("Illegal GreetRequest: %v\n", err)
		return nil, err /*myerr.GRError{Code: 5000, Message: "greet error"}*/
	}

	// Print GreetRequest
	s.printGreetRequest(in)
	//return &pb.GreetReply{Ip: &ip, Message: in.GetIp().GetName() + "! Hello! Connect successfully!"}, nil
	return &pb.GreetReply{Ip: &ip, Message: "Hello"}, nil
}

func (s *server) Login(ctx context.Context, in *pb.AccessRequest) (*pb.AccessReply, error) {
	// Validate AccessRequest
	if err := in.Validate(); err != nil {
		// If validation failed, server should return error message.
		log.Printf("Fail to login: %v\n", err)
		return nil, myerr.LRError{Code: 5000, Message: "login error"}
	}

	// Print AccessRequest
	s.printAccessRequest(in)
	return &pb.AccessReply{Ip: &ip, Message: "FeedBack: LOGIN SUCCESS!"}, nil
}

// Print GreetRequest
func (s *server) printGreetRequest(in *pb.GreetRequest) {
	log.Println("~~~~~~~~~ Receive Greet Request ~~~~~~~~~")
	/*
		fmt.Println("Request from: ", in.GetIp().GetAddr(), ":", in.GetIp().GetPort(), in.GetIp().GetName())
		fmt.Println("ID: ", in.GetId())
		fmt.Println("Message: ", in.GetMessage())
		fmt.Println("Num: wheel.Double: ", in.Num.Value)
	*/
	fmt.Println("doublenum: ", in.Doublenum.Value)
	fmt.Println("floatnum: ", in.Floatnum.Value)
	fmt.Println("int64num: ", in.Int64Num.Value)
	fmt.Println("uint64num: ", in.Uint64Num.Value)
	fmt.Println("int32num: ", in.Int32Num.Value)
	fmt.Println("uint32num: ", in.Uint32Num.Value)
	fmt.Println("str: ", in.Str.Value)
	fmt.Println("bytes: ", string(in.B.Value))
	fmt.Println("bool: ", in.Bo.Value)
}

// Print AccessRequest
func (s *server) printAccessRequest(in *pb.AccessRequest) {
	log.Println("~~~~~~~~~ Receive Login Access Request ~~~~~~~~~")
	fmt.Println("Request from: ", in.GetIp().GetAddr(), ":", in.GetIp().GetPort(), in.GetIp().GetName())
	fmt.Println("Email: ", in.GetEmail().GetEmail())
	fmt.Println("Chinese: ", in.GetIsCN())
}

func main() {
	fmt.Println("=============== Here is", ip.GetName(), "(", ip.GetAddr(), ":", ip.GetPort(), ") ===============")

	lis, err := net.Listen("tcp", "localhost:6666") // Listen localhost port, tcp
	if err != nil {
		log.Fatalf("Fail to listen port: %v", err)
	}

	s := grpc.NewServer() // Create a gRPC server which has no service registered and has not started to accept requests yet.
	// Register a service and its implementation to the gRPC server.
	pb.RegisterCommunicationServer(s, &server{})

	reflection.Register(s) // Register the server reflection service on the given gRPC server.
	// Serve accepts incoming connections on the listener lis, creating a new
	// ServerTransport and service goroutine for each. The service goroutines
	// read gRPC requests and then call the registered handlers to reply to them.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
