package main

import (
	"fmt"
	pb "github.com/hzxiao/microservice-tutorial/consignment-service/proto/consignment"
	vesselProto "github.com/hzxiao/microservice-tutorial/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
	"os"
)


const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(micro.Name("go.micro.srv.consignment"))
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	h := &handler{repository, vesselClient}
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
