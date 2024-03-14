package repository

import (
	"context"

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
	CreateSchool(ctx context.Context)
	ReadSchool(ctx context.Context)
	UpdateSchool(ctx context.Context)
	DeleteSchool(ctx context.Context)
	NewPassword(ctx context.Context)
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
