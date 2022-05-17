package main

import (
	"context"
	pb "grpc3/calculator/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	// doCalculator(c)
	doPrimes(c)
}

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{
		FirstNumber: 12390392840,
	})

	if err != nil {
		log.Fatalf("Error while calling primes: %v\n", err)
	}

	// log.Printf("Primes: %d\n", res.Result)

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v \n", err)
		}

		log.Printf("Primes: %d\n", msg.Result)
	}
}
