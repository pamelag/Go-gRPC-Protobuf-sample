package main

import (
	"log"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../graph_proto"
)

const (
	address     = "localhost:50052"
	defaultName = "DocCluster"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGraphClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.CreateGraph(context.Background(), &pb.GraphRequest{Name: name})
	if err != nil {
		log.Fatalf("could not create graph : %v", err)
	}
	log.Printf("Greeting: %s", r.Message)


}
