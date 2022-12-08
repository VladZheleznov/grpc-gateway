package main

import (
	"context"
	pb "grpc-gateway/gen/proto"
	"log"
	"net"
	"net/http"

	proto_user "github.com/AlexeyBorisovets/USER/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
	uc proto_user.USERClient
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Registr(ctx context.Context, req *pb.RegistrRequest) (*pb.RegistrResponse, error) {
	test, err := s.uc.Registr(ctx, &proto_user.RegistrRequest{})
	if err != nil{
		log.Fatalln(err)
	}
	return &proto_user.RegistrResponse{}, nil
}

func main() {

	go func() {
		mux := runtime.NewServeMux()

		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})

		log.Fatalln(http.ListenAndServe("localhost:50052", mux))
	}()

	lestner, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	grpcServer.Serve(lestner)
	if err != nil {
		log.Println(err)
	}

}
