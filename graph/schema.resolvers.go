package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alai04/gql-truck/graph/generated"
	"github.com/alai04/gql-truck/graph/model"
	"github.com/alai04/gql-truck/pkg/cars"
)

func (r *queryResolver) ApprovedCars(ctx context.Context) ([]*model.TruckCar, error) {
	// panic(fmt.Errorf("not implemented"))
	return cars.ApprovedCars(), nil
}

func (r *queryResolver) Approved(ctx context.Context, input model.Plate) (bool, error) {
	// panic(fmt.Errorf("not implemented"))
	return cars.Approved(input.Plate), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
