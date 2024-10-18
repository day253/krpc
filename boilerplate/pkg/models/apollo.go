package models

import (
	"context"

	"github.com/samber/do"
)

const (
	ApolloName = "apollo"
	ApolloType = "apollo"
)

type Apollo struct {
	hooks DetailsHooksChain
}

func (p *Apollo) Name() string {
	return ApolloName
}

func (p *Apollo) Type() string {
	return ApolloType
}

func (p *Apollo) Clone() (Model, error) {
	return &Apollo{
		hooks: p.hooks,
	}, nil
}

func (p *Apollo) Run(ctx context.Context, input ModelInput) (ModelResult, error) {
	return NewMapModelResult(
		p.hooks.Run(
			ctx,
			input.Input(),
		),
	), nil
}

func NewApollo(i *do.Injector) (Model, error) {
	return &Apollo{
		hooks: DetailsHooksChain{
			&Move{},
		},
	}, nil
}

func init() {
	do.ProvideNamed(Injector, ApolloName, NewApollo)
}
