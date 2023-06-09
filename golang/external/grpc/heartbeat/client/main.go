package main

import (
	"context"
	pb "golang/external/grpc/heartbeat/protobuf"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:13335", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHerBeaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	/* ECHO */
	r, err := c.Echo(ctx, &pb.EchoRequest{SendMessage: "idiot"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Get Echo Message: %s", r.GetEchoMessage())

	/* STATUS */
	// r, err := c.Status(ctx, &pb.StatusRequest{Sender: "john"})
	// if err != nil {
	// 	log.Fatalf("could not get status: %v", err)
	// }
	// // fmt.Println(r.ProtoReflect().Descriptor())
	// jsonStr, _ := json.MarshalIndent(r, "", "  ")
	// log.Printf("Get Status From Server: %v", string(jsonStr))

	ctx2, ctxCancelFunc := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	connection, err := grpc.DialContext(ctx, "127.0.0.1:13335",
		grpc.WithBlock(),
		grpc.WithDisableRetry(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ctxCancelFunc()
		return
	}
	client := pb.NewHerBeaClient(connection)

	/* ECHO */
	r, err = client.Echo(ctx2, &pb.EchoRequest{SendMessage: "omg"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Get Echo Message: %s", r.GetEchoMessage())
}
