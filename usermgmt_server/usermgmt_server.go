package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct{
	pb.UnimplementedUserManagementServer
	userList *pb.UserList
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		userList: &pb.UserList{},
	}
}

func (server *UserManagementServer) Run() error {
	lis , err := net.Listen("tcp" , port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s , server)
	log.Printf("server listening at %v address\n" , lis.Addr() )
	return s.Serve(lis) 
}

func (s *UserManagementServer) CreateNewUser( ctx context.Context , in *pb.NewUser ) (*pb.User , error) {
	log.Printf("Received %v " , in.GetName())
	var user_id int32 = int32(rand.Int31n(1000))
	createdUser := &pb.User{ Name: in.GetName() , Age: in.GetAge() , Id: user_id , City: in.City , Phone: in.Phone , Height: in.Height , Married: in.Married }
	s.userList.Users = append(s.userList.Users, createdUser)
	return  createdUser, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context , in *pb.GetUsersParams)(*pb.UserList , error){
	return s.userList , nil
}

func main(){
	var userManagementServer *UserManagementServer = NewUserManagementServer()
	if err := userManagementServer.Run() ; err != nil {
		log.Fatal(err)
	}
}