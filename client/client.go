package main

import (
	"context"
	"go-grpc/chat"
	"log"

	"google.golang.org/grpc"
)

func main(){
	var conn *grpc.ClientConn
	conn,err:=grpc.Dial(":50005",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("could not connect port 50005 : %v",err)
	}
	defer conn.Close()
	c:=chat.NewChatSeviceClient(conn)
	message:=chat.Message{
		Body: "helloo from client..",
	}

	response,err:=c.SayHello(context.Background(),&message)
	if err!=nil{
		log.Fatalf("error when calling sayhello : %v",err)
	}
	log.Printf("response from server :%s",response.Body)
}