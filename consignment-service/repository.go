package main

import (
	"context"
	"log"

	pb "github.com/knightdave/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	log.Printf("Getting consignments")

	cur, err := repository.collection.Find(context.Background(), bson.D{}, )
	if err != nil {
		return nil, err
	}
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			log.Printf("Log Error %v", err)
			return nil, err
		}
		consignments = append(consignments, consignment)
	}

	return consignments, err
}
