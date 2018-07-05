package main

import (
	"context"
	pb "example/communication"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// Define localhost address infomation
const (
	dstaddr     = "127.0.0.1:6666"
	address     = "127.0.0.1"
	defaultname = "client2"
	port        = 9999
)

// set ip
var ip = pb.IP{
	Addr: address,
	Name: defaultname,
	Port: port,
}

// Print greet-reply
func printGreetReply(out *pb.GreetReply) {
	fmt.Println("~~~~~~~~~ Receive Greet Reply ~~~~~~~~~")
	fmt.Println("Reply from: ", out.Ip.GetAddr(), ":", out.Ip.GetPort(), out.Ip.GetName())
	fmt.Println("Message: ", out.GetMessage())
}

// Print login-reply
func printAccessReply(out *pb.AccessReply) {
	fmt.Println("~~~~~~~~~ Receive Access Reply ~~~~~~~~~")
	fmt.Println("Reply from: ", out.Ip.GetAddr(), ":", out.Ip.GetPort(), out.Ip.GetName())
	fmt.Println("Message: ", out.GetMessage())
}

func main() {
	// Set up a connection with the server.
	conn, err := grpc.Dial(dstaddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect: %v", err)
	}

	// Avoid memory leak.
	defer conn.Close()

	// Create a client.
	c := pb.NewCommunicationClient(conn)

	// User defined request
	var sendgreet = flag.Bool("sendgreet", false, "(Optional) Whether send greet request or not. -sendgreet=true") // whether send greet request or not
	var sendlogin = flag.Bool("sendlogin", false, "(Optional) Whether send login request or not. -sendlogin=true") // whether send login request or not

	// Declare the input greetrequest
	var id = flag.Uint64("id", 10000, "ID of client. This is valid only when sendgreet field is set true.")
	var greetmessage = flag.String("greetmessage", "", "Greet request message. This is valid only when sendgreet field is set true.")
	// Declare the input accessreply
	var iscn = flag.Bool("iscn", false, "Whether is China or not. This is valid only when sendlogin field is set true.")
	var email = flag.String("email", "userexample@qq.com", "Email address. This is valid only when sendlogin field is set true.")
	var pwd = flag.String("pwd", "user", "Email password. This is valid only when sendaccess field is set true.")
	// Parse input
	flag.Parse()

	fmt.Println("=============== Here is", ip.GetName(), "(", ip.GetAddr(), ":", ip.GetPort(), ") ===============")

	// Send greet request.
	if *sendgreet {
		// Greet with the server.
		log.Println("********** Send Greet Request **********")
		greetreply, err := c.Greet(context.Background(), &pb.GreetRequest{Id: *id, Ip: &ip, Message: *greetmessage})
		// Request failed
		if greetreply == nil && err != nil {
			log.Fatalf("Please check if the input is valid!%v", err)
		}

		// Validate GreetReply
		if err = greetreply.Validate(); err != nil {
			log.Fatalf("GreetReply validate error: %v", err)
		}

		// Print out the reply from server.
		printGreetReply(greetreply)
	}

	// Send login request.
	if *sendlogin {
		// SendMessage to server.
		log.Println("********** Send Login Request **********")
		accessreply, err := c.Login(context.Background(), &pb.AccessRequest{Ip: &ip, IsCN: *iscn, Email: &pb.AccessRequest_Email{Email: *email, Pwd: *pwd}})
		// Request failed
		if accessreply == nil && err != nil {
			log.Fatalf("Please check if the input is valid!\n%v", err)
		}

		// Validate AccessReply
		if err = accessreply.Validate(); err != nil {
			log.Fatalf("AccessReply validate error: %v", err)
		}
		// Print out the reply from server.
		printAccessReply(accessreply)
	}
}
