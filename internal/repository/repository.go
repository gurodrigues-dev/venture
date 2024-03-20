package repository

import (
	"context"
	"gin/types"

	_ "github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context)
	ReadUser(ctx context.Context)
	UpdateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
	CreateChild(ctx context.Context)
	ReadChild(ctx context.Context)
	UpdateChild(ctx context.Context)
	DeleteChild(ctx context.Context)
	CreateDriver(ctx context.Context)
	ReadDriver(ctx context.Context)
	UpdateDriver(ctx context.Context)
	DeleteDriver(ctx context.Context)
	CreateSchool(ctx context.Context, school *types.School) error
	ReadSchool(ctx context.Context, id *int) (*types.School, error)
	UpdateSchool(ctx context.Context) error
	DeleteSchool(ctx context.Context, id *int) error
	AuthSchool(ctx context.Context, school *types.School) error
	NewPassword(ctx context.Context)
	VerifyEmailExists(ctx context.Context, email *string) (bool, error)
}

type CloudRepository interface {
	CheckEmail(ctx context.Context)
	SendEmail(ctx context.Context)
	SaveImageBucket(ctx context.Context)
}

type CacheRepository interface {
	SaveKeyAndValue(ctx context.Context)
	VerifyToken(ctx context.Context)
}
