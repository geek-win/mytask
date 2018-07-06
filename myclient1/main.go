package main

import (
	"context"
	pb "example/communication"
	"flag"
	"fmt"
	"log"
	"wtypes"

	"google.golang.org/grpc"
)

// Define localhost address infomation
const (
	dstaddr     = "127.0.0.1:6666"
	address     = "127.0.0.1"
	defaultname = "client1"
	port        = 8888
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

	/*
		// Declare the input greetrequest
		var id = flag.Uint64("id", 10000, "ID of client. This is valid only when sendgreet field is set true.")
		var greetmessage = flag.String("greetmessage", "", "Greet request message. This is valid only when sendgreet field is set true.")
	*/
	// Declare the input accessreply
	var iscn = flag.Bool("iscn", false, "Whether is China or not. This is valid only when sendlogin field is set true.")
	var email = flag.String("email", "userexample@qq.com", "Email address. This is valid only when sendlogin field is set true.")
	var pwd = flag.String("pwd", "user", "Email password. This is valid only when sendaccess field is set true.")

	var floatv = flag.Bool("floatv", false, "input wtypes.Float.valid")
	var floatn = flag.Float64("floatn", 0, "input wtypes.Float.val")
	var doublev = flag.Bool("doublev", false, "input wtypes.Double.valid")
	var doublen = flag.Float64("doublen", 0, "input wtypes.Double.val")
	var int64v = flag.Bool("int64v", false, "input wtypes.Int64.valid")
	var int64n = flag.Int64("int64n", 0, "input wtypes.Int64.val")
	var uint64v = flag.Bool("uint64v", false, "input wtypes.UInt64.valid")
	var uint64n = flag.Uint64("uint64n", 0, "input wtypes.UInt64.val")
	var int32v = flag.Bool("int32v", false, "input wtypes.Int32.valid")
	var int32n = flag.Int("int32n", 0, "input wtypes.Int32.val")
	var uint32v = flag.Bool("uint32v", false, "input wtypes.Uint32.valid")
	var uint32n = flag.Uint("uint32n", 0, "input wtypes.Uint32.val")
	var strv = flag.Bool("strv", false, "input wtypes.String.valid")
	var strn = flag.String("strn", "", "input wtypes.String.val")
	var bov = flag.Bool("bov", false, "input wtypes.Bool.valid")
	var bon = flag.Bool("bon", false, "input wtypes.Bool.val")
	var bytesv = flag.Bool("bytesv", false, "input wtypes.Bytes.valid")
	var bytesn = flag.String("bytesn", "", "input wtypes.Bytes.val")
	// Parse input
	flag.Parse()

	fmt.Println("=============== Here is", ip.GetName(), "(", ip.GetAddr(), ":", ip.GetPort(), ") ===============")

	// Send greet request.
	if *sendgreet {
		// Greet with the server.
		log.Println("********** Send Greet Request **********")
		//greetreply, err := c.Greet(context.Background(), &pb.GreetRequest{Id: *id, Ip: &ip, Message: *greetmessage, Num: &wtypes.Double{Valid: *isvalid, Value: *value}})
		greetreply, err := c.Greet(context.Background(), &pb.GreetRequest{Doublenum: &wtypes.Double{Valid: *doublev, Value: *doublen}, Floatnum: &wtypes.Float{Valid: *floatv, Value: float32(*floatn)}, Int64Num: &wtypes.Int64{Valid: *int64v, Value: *int64n}, Uint64Num: &wtypes.UInt64{Valid: *uint64v, Value: *uint64n}, Int32Num: &wtypes.Int32{Valid: *int32v, Value: int32(*int32n)}, Uint32Num: &wtypes.UInt32{Valid: *uint32v, Value: uint32(*uint32n)}, Str: &wtypes.String{Valid: *strv, Value: *strn}, Bo: &wtypes.Bool{Valid: *bov, Value: *bon}, B: &wtypes.Bytes{Valid: *bytesv, Value: []byte(*bytesn)}})
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
