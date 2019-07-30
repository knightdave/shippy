package main

import (
	"context"
	"errors"
	pb "github.com/knightdave/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"

)

// Repository - 
type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

// MongoRepository - 
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(vessel *pb.Vessel) error {
	_, err := repository.collection.InsertOne(context.Background(), vessel)
	return err
}


// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *MongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var vessel *pb.Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}

	}
	return nil, errors.New("No vessel found by that spec")
}
