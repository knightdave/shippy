package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/knightdave/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

func createDummyData(repo repository) {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		err := repo.Create(v)
		if err != nil {
			log.Fatalf("Error in creating DummyData in vessel")
		}
	}
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient("mongodb://" + uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{
		vesselCollection,
	}
	createDummyData(repository)

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})


	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
