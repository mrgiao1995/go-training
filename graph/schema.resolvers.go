package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-training/graph/generated"
	"go-training/graph/model"
)

// ChangeCustomerPassword is the resolver for the changeCustomerPassword field.
func (r *mutationResolver) ChangeCustomerPassword(ctx context.Context, input model.ChangeCustomerPassword) (*model.Empty, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateNewCustomer) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateCustomer is the resolver for the updateCustomer field.
func (r *mutationResolver) UpdateCustomer(ctx context.Context, input model.UpdateCustomer) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// CustomerBooking is the resolver for the customerBooking field.
func (r *queryResolver) CustomerBooking(ctx context.Context, input model.ViewCustomerBookingHistories) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Customer(ctx context.Context) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}
