package repository

import (
	"context"
	"gin/types"

	_ "github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context, user *types.User) error
	ReadUser(ctx context.Context, id *int) (*types.User, error)
	UpdateUser(ctx context.Context) error
	DeleteUser(ctx context.Context, cpf *string) error
	AuthUser(ctx context.Context, user *types.User) (*types.User, error)
	CreateChild(ctx context.Context, child *types.Child, user *types.User) error
	ReadChild(ctx context.Context, idUser *int) ([]types.Child, error)
	UpdateChild(ctx context.Context) error
	DeleteChild(ctx context.Context, idChild *int) error
	CreateDriver(ctx context.Context, driver *types.Driver) error
	ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error)
	UpdateDriver(ctx context.Context) error
	DeleteDriver(ctx context.Context, cnh *string) error
	AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error)
	CreateSchool(ctx context.Context, school *types.School) error
	ReadSchool(ctx context.Context, cnpj *string) (*types.School, error)
	ReadAllSchools(ctx context.Context) ([]types.School, error)
	UpdateSchool(ctx context.Context) error
	DeleteSchool(ctx context.Context, cnpj *string) error
	AuthSchool(ctx context.Context, school *types.School) (*types.School, error)
	NewPassword(ctx context.Context)
	VerifyEmailExists(ctx context.Context, table, email *string) (bool, error)
	CreateInvite(ctx context.Context, invite *types.Invite) error
	ReadInvite(ctx context.Context, invite_id *int) (*types.Invite, error)
	ReadAllInvites(ctx context.Context, cnh *string) ([]types.Invite, error)
	UpdateInvite(ctx context.Context, invite_id *int) error
	DeleteInvite(ctx context.Context, invite_id *int) error
	CreateRecordToSchoolAndDriver(ctx context.Context, invite *types.Invite) error
	DeleteRecordToSchoolAndDriver(ctx context.Context, record_id *int) error
	GetDriversSchool(ctx context.Context, cnh *string) ([]types.School, error)
}

type Cloud interface {
	CreateAndSaveQrCodeInS3(ctx context.Context, cnh *string) (string, error)
}

type Cache interface {
	SaveKeyAndValue(ctx context.Context)
	FindKeyRedis(ctx context.Context)
	ValidateExistsSismember(ctx context.Context)
}

type Messaging interface {
	Producer(ctx context.Context, msg string) error
}
