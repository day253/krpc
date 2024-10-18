package models

import (
	"context"
)

const (
	PlaceholderType = "placeholder"
)

type Placeholder struct {
	name  string
	hooks DetailsHooksChain
}

func (p *Placeholder) Name() string {
	return p.name
}

func (p *Placeholder) Type() string {
	return PlaceholderType
}

func (p *Placeholder) Clone() (Model, error) {
	return &Placeholder{
		name:  p.name,
		hooks: p.hooks,
	}, nil
}

func (p *Placeholder) Run(ctx context.Context, input ModelInput) (ModelResult, error) {
	return NewMapModelResult(
		p.hooks.Run(
			ctx,
			input.Input(),
		),
	), nil
}

func NewPlaceholder(name string) (Model, error) {
	move := &Move{}
	return &Placeholder{
		name: name,
		hooks: DetailsHooksChain{
			move,
		},
	}, nil
}

func GetPlaceholder(name string) (Model, error) {
	return NewPlaceholder(name) // 每次都是新的不需要拷贝
}
