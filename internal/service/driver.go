package service

import (
	"context"
	"fmt"
	"gin/config"
	"gin/internal/repository"
	"gin/types"
	"gin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func (ds *DriverService) CreateEmployee(ctx context.Context, invite *types.Invite) error {
	return ds.driverrepository.CreateEmployee(ctx, invite)
}

func (ds *DriverService) ParserJwtDriver(ctx *gin.Context) (interface{}, error) {

	cnh, found := ctx.Get("cnh")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnh, nil

}

func (ds *DriverService) CreateTokenJWTDriver(ctx context.Context, driver *types.Driver) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cnh": driver.CNH,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil
}
