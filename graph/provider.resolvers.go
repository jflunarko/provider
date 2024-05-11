package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"myapp/graph/generated"
	"myapp/model"
	"myapp/service"
)

// Packet is the resolver for the packet field.
func (r *providerResolver) Packet(ctx context.Context, obj *model.Provider) ([]*model.Packet, error) {
	s := service.GetService()

	packet, _ := s.PacketGetByProviderId(ctx, obj.ID)
	return packet, nil
}

// Create is the resolver for the create field.
func (r *providerMutationsResolver) Create(ctx context.Context, obj *model.ProviderMutations, input model.NewProvider) (*model.Provider, error) {
	s := service.GetTransaction()

	provider, _ := s.CreateProvider(ctx, input)
	s.Commit()
	return provider, nil
}

// Delete is the resolver for the delete field.
func (r *providerMutationsResolver) Delete(ctx context.Context, obj *model.ProviderMutations, id int) (string, error) {
	s := service.GetTransaction()

	provider, _ := s.DeleteProvider(ctx, id)
	s.Commit()
	return provider, nil
}

// Update is the resolver for the update field.
func (r *providerMutationsResolver) Update(ctx context.Context, obj *model.ProviderMutations, id int, input model.UpdateProvider) (*model.Provider, error) {
	s := service.GetTransaction()

	provider, _ := s.UpdateProvider(ctx, input, id)
	s.Commit()
	return provider, nil
}

// Providers is the resolver for the providers field.
func (r *providerQueryResolver) Providers(ctx context.Context, obj *model.ProviderQuery) ([]*model.Provider, error) {
	s := service.GetService()

	providers, _ := s.ProvidersGetAll(ctx)
	return providers, nil
}

// Provider is the resolver for the provider field.
func (r *providerQueryResolver) Provider(ctx context.Context, obj *model.ProviderQuery, id int) (*model.Provider, error) {
	s := service.GetService()
	provider, _ := s.ProviderGetById(ctx, id)
	return provider, nil
}

// Provider returns generated.ProviderResolver implementation.
func (r *Resolver) Provider() generated.ProviderResolver { return &providerResolver{r} }

// ProviderMutations returns generated.ProviderMutationsResolver implementation.
func (r *Resolver) ProviderMutations() generated.ProviderMutationsResolver {
	return &providerMutationsResolver{r}
}

// ProviderQuery returns generated.ProviderQueryResolver implementation.
func (r *Resolver) ProviderQuery() generated.ProviderQueryResolver { return &providerQueryResolver{r} }

type providerResolver struct{ *Resolver }
type providerMutationsResolver struct{ *Resolver }
type providerQueryResolver struct{ *Resolver }
