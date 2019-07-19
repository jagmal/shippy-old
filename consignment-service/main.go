package main

import (
	"context"
	pb "github.com/jagmal/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/jagmal/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("shippy.service.consignment")
	)

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

	repository := &MongoRepository(consignmentCollection)
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.client", srv.Client())
	h := &handler(repository, vesselClient)

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf(err)
	}
}
