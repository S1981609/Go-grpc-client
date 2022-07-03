package main

import (
	generatedfiles "Go-grpc-client/generatedfiles"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr       = flag.String("addr", "localhost:8056", "the address to connect to")
	name       = flag.String("Name", "George", "Name to greet")
	coursecode = flag.String("coursecode", "CS12343", "course code details")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := generatedfiles.NewAssignTeamMateClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AssignTeamMate(ctx, &generatedfiles.AssignTeamRequestgRPC{Name: *name, Id: 113527, CourseCode: *coursecode})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %v", r)
}
