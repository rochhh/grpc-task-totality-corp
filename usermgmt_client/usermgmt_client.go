package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type NewUser struct{
	pb.NewUser
}

type User struct{
	pb.User
}

func (nu *NewUser) NewUserConstructor() *User{
	return &User{}
}

func main(){ 
	conn , err := grpc.Dial(address,grpc.WithInsecure() , grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)
	ctx , cancel := context.WithTimeout(context.Background() , time.Second)
	defer cancel()
	var newUsers = make(map[string]*pb.User)
	newUsers[ "key" ] = &pb.User{
		Name : "Foo",
		Age : 22,
		City : "Mumbai",
		Phone : 1234,
		Height: 5.9,
		Married : false,
	}
	params := &pb.GetUsersParams{}
	r , err := c.GetUsers(ctx , params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("user list")
	fmt.Printf("the users %v\n" , r.GetUsers())
}