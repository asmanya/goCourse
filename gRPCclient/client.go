package main

import (
	"context"
	"log"
	"time"

	mainapipb "simplegrpcclient/proto/gen"
	pb "simplegrpcclient/proto/gen"
	fw "simplegrpcclient/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Failed to load certificate", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := pb.NewCalculateClient(conn)
	client2 := pb.NewGreeterClient(conn)
	// fwClient := fw.NewAufWiedersehenClient(conn)
	client3 := mainapipb.NewBidFarewellClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.AddRequest{
		A: 10,
		B: 20,
	}

	// =============
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	// ==============
	reqGreet := &pb.HelloRequest{
		Name: "John",
	}
	res1, err := client2.Greet(ctx, reqGreet)
	if err != nil {
		log.Fatalln("Could not Greet", err)
	}

	resAddFromGreeter, err := client2.Add(ctx, reqGreet)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	// ===============
	reqGoodbye := &fw.GoodByeRequest{
		Name: "Jane",
	}
	// resFw, err := fwClient.BidGoodBye(ctx, reqGoodbye)
	// if err != nil {
	// 	log.Fatalln("Could not bid goodbye", err)
	// }

	res3, err := client3.BidGoodBye(ctx, reqGoodbye)
	if err != nil {
		log.Fatalln("Could not bid goodbye", err)
	}

	log.Println("Sum:", res.Sum)
	log.Println("Greeting message:", res1.Message)
	log.Println("+++++++++++++++Greeting message from second add function in proto file:", resAddFromGreeter.Message)
	log.Println("Bid goodbye:", res3.Message)

	// log.Println("Bid goodbye:", resFw.Message)
	// state := conn.GetState()
	// log.Println("Connection State:", state)

}
