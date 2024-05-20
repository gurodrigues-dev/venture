package repository

import (
	"context"
	"gin/types"
)

type DriverRepository interface {
	CreateDriver(ctx context.Context, driver *types.Driver) error
	ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error)
	UpdateDriver(ctx context.Context) error
	DeleteDriver(ctx context.Context, cnh *string) error
	AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error)
	ReadInvite(ctx context.Context, invite_id *int) (*types.Invite, error)
	ReadAllInvites(ctx context.Context, cnh *string) ([]types.Invite, error)
	UpdateInvite(ctx context.Context, invite_id *int) error
	DeleteInvite(ctx context.Context, invite_id *int) error
	GetWorkplaces(ctx context.Context, cnh *string) ([]types.School, error)
}
