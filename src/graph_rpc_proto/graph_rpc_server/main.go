package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../graph_proto"
	"google.golang.org/grpc/reflection"
	graph "../graph"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) CreateGraph(ctx context.Context, in *pb.GraphRequest) (*pb.GraphReply, error) {
	graph := graph.NewGraphStruct()
	graph.SetName(in.Name)

	//theGraph := graph.NewGraphStruct()
	//theGraph.AddVertex('A'); // 0 (start for dfs)
	//theGraph.addVertex(‘B’); // 1
	//theGraph.addVertex(‘C’); // 2
	//theGraph.addVertex(‘D’); // 3
	//theGraph.addVertex(‘E’); // 4
	//theGraph.addEdge(0, 1); // AB
	//theGraph.addEdge(1, 2); // BC
	//theGraph.addEdge(0, 3); // AD
	//theGraph.addEdge(3, 4); // DE
	//System.out.print("Visits: ");
	//theGraph.dfs() // depth-first search
	//System.out.println()

	return &pb.GraphReply{Message: "Create Graph " + graph.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGraphServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
