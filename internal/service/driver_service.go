package service

import (
	"context"
	"gin/internal/repository"
	"gin/types"
	"gin/utils"
)

type DriverService struct {
	driverrepository repository.DriverRepository
}

func NewDriverService(repo repository.DriverRepository) *DriverService {
	return &DriverService{
		driverrepository: repo,
	}
}

func (ds *DriverService) CreateDriver(ctx context.Context, driver *types.Driver) error {
	driver.Password = utils.HashPassword(driver.Password)
	return ds.driverrepository.CreateDriver(ctx, driver)
}

func (ds *DriverService) ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error) {
	return ds.driverrepository.ReadDriver(ctx, cnh)
}

func (ds *DriverService) UpdateDriver(ctx context.Context) error {
	return ds.driverrepository.UpdateDriver(ctx)
}

func (ds *DriverService) DeleteDriver(ctx context.Context, cnh *string) error {
	return ds.driverrepository.DeleteDriver(ctx, cnh)
}

func (ds *DriverService) AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error) {
	driver.Password = utils.HashPassword(driver.Password)
	return ds.driverrepository.AuthDriver(ctx, driver)
}

func (ds *DriverService) ReadInvite(ctx context.Context, invite_id *int) (*types.Invite, error) {
	return ds.driverrepository.ReadInvite(ctx, invite_id)
}

func (ds *DriverService) ReadAllInvites(ctx context.Context, cnh *string) ([]types.Invite, error) {
	return ds.driverrepository.ReadAllInvites(ctx, cnh)
}

func (ds *DriverService) UpdateInvite(ctx context.Context, invite_id *int) error {
	return ds.driverrepository.UpdateInvite(ctx, invite_id)
}

func (ds *DriverService) DeleteInvite(ctx context.Context, invite_id *int) error {
	return ds.driverrepository.DeleteInvite(ctx, invite_id)
}

func (ds *DriverService) GetWorkplaces(ctx context.Context, cnh *string) ([]types.School, error) {
	return ds.driverrepository.GetWorkplaces(ctx, cnh)
}
