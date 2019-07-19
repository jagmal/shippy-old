package main

import (
	"context"
	pb "github.com/jagmal/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository --
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// MongoVesselRepository -
type MongoVesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable = checks a specification against a map of vessels.
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *MongoVesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// Note: we can alternatively put query to filter out matching vessels only
	cur, err := repo.collection.Find(context.Background(), nil, nil)

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
