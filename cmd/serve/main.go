package main

import (
    "log"
    "net"
    "fmt"

    chord "chord/pkg"
    chordv1 "chord/pkg/v1"
    chordpbv1 "chord/proto/v1"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {

    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
    if err != nil {
        log.Fatal(err.Error())
    }

    service := chordv1.NewNodeService(chord.NewNode(
        "127.0.0.1", "8080", 10, chord.NewSimpleHasher(),
    ))

    grpcServer := grpc.NewServer()
    chordpbv1.RegisterNodeServiceServer(grpcServer, service)
    reflection.Register(grpcServer)

    err = grpcServer.Serve(lis)
    if err != nil {
        log.Fatal(err.Error())
    }
}
