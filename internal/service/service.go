package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gin/config"
	"gin/internal/repository"
	"gin/types"
	"gin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repository repository.Repository
	cloud      repository.Cloud
	redis      repository.Cache
	broker     repository.Messaging
}

func New(repo repository.Repository, cloud repository.Cloud, redis repository.Cache, broker repository.Messaging) *Service {
	return &Service{
		repository: repo,
		cloud:      cloud,
		redis:      redis,
		broker:     broker,
	}
}

func (s *Service) CreateUser(ctx context.Context) {

}

func (s *Service) ReadUser(ctx context.Context) {

}

func (s *Service) UpdateUser(ctx context.Context) {

}

func (s *Service) DeleteUser(ctx context.Context) {

}

func (s *Service) CreateChild(ctx context.Context) {

}

func (s *Service) ReadChild(ctx context.Context) {

}

func (s *Service) UpdateChild(ctx context.Context) {

}

func (s *Service) DeleteChild(ctx context.Context) {

}

func (s *Service) CreateDriver(ctx context.Context, driver *types.Driver) error {
	driver.Password = utils.HashPassword(driver.Password)
	return s.repository.CreateDriver(ctx, driver)
}

func (s *Service) ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error) {
	return s.repository.ReadDriver(ctx, cnh)
}

func (s *Service) UpdateDriver(ctx context.Context) error {
	return s.repository.UpdateDriver(ctx)
}

func (s *Service) DeleteDriver(ctx context.Context, cnh *string) error {
	return s.repository.DeleteDriver(ctx, cnh)
}

func (s *Service) AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error) {
	driver.Password = utils.HashPassword(driver.Password)
	return s.repository.AuthDriver(ctx, driver)
}

func (s *Service) CreateSchool(ctx context.Context, school *types.School) error {
	school.Password = utils.HashPassword(school.Password)
	return s.repository.CreateSchool(ctx, school)
}

func (s *Service) ReadSchool(ctx context.Context, cnpj *string) (*types.School, error) {
	return s.repository.ReadSchool(ctx, cnpj)
}

func (s *Service) ReadAllSchools(ctx context.Context) ([]types.School, error) {
	return s.repository.ReadAllSchools(ctx)
}

func (s *Service) UpdateCreateSchool(ctx context.Context) error {
	return s.repository.UpdateSchool(ctx)
}

func (s *Service) DeleteSchool(ctx context.Context, cnpj *string) error {
	return s.repository.DeleteSchool(ctx, cnpj)
}

func (s *Service) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	school.Password = utils.HashPassword(school.Password)
	return s.repository.AuthSchool(ctx, school)
}

func (s *Service) AddMessageInQueue(ctx context.Context, msg string) error {
	return s.broker.Producer(ctx, msg)
}

func (s *Service) CreateAndSaveQrCodeInS3(ctx context.Context, cnh *string) (string, error) {
	return s.cloud.CreateAndSaveQrCodeInS3(ctx, cnh)
}

func (s *Service) SaveKeyAndValue(ctx context.Context) {

}

func (s *Service) VerifyToken(ctx context.Context) {

}

func (s *Service) CreateTokenJWTUser(ctx context.Context, user *types.User) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cpf": user.CPF,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *Service) CreateTokenJWTDriver(ctx context.Context, driver *types.Driver) (string, error) {

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

func (s *Service) CreateTokenJWTSchool(ctx context.Context, school *types.School) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cnpj": school.CNPJ,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil

}

func (s *Service) ParserJwtSchool(ctx *gin.Context) (interface{}, error) {

	cnpj, found := ctx.Get("cnpj")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnpj, nil

}

func (s *Service) ParserJwtUser(ctx *gin.Context) (interface{}, error) {

	cpf, found := ctx.Get("cpf")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cpf, nil

}

func (s *Service) ParserJwtDriver(ctx *gin.Context) (interface{}, error) {

	cnh, found := ctx.Get("cnh")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnh, nil

}

func (s *Service) InterfaceToString(value interface{}) (*string, error) {
	switch v := value.(type) {
	case string:
		return &v, nil
	default:
		return nil, fmt.Errorf("value isn't string")
	}
}

func (s *Service) EmailStructToJSON(email *types.Email) (string, error) {

	json, err := json.Marshal(email)

	if err != nil {
		return "", err
	}

	return string(json), nil

}

func (s *Service) CreateInvite(ctx context.Context, invite *types.Invite) error {
	return s.repository.CreateInvite(ctx, invite)
}

func (s *Service) ReadInvite(ctx context.Context, invite_id *int) (*types.Invite, error) {
	return s.repository.ReadInvite(ctx, invite_id)
}

func (s *Service) ReadAllInvites(ctx context.Context, cnh *string) ([]types.Invite, error) {
	return s.repository.ReadAllInvites(ctx, cnh)
}

func (s *Service) UpdateInvite(ctx context.Context, invite_id *int) error {
	return s.repository.UpdateInvite(ctx, invite_id)
}

func (s *Service) DeleteInvite(ctx context.Context, invite_id *int) error {
	return s.repository.DeleteInvite(ctx, invite_id)
}

func (s *Service) CreateRecordToSchoolAndDriver(ctx context.Context, invite *types.Invite) error {
	return s.repository.CreateRecordToSchoolAndDriver(ctx, invite)
}

func (s *Service) GetWorkplaces(ctx context.Context, cnh *string) ([]types.School, error) {
	return s.repository.GetWorkplaces(ctx, cnh)
}

func (s *Service) GetEmployees(ctx context.Context, cnpj *string) ([]types.Driver, error) {
	return s.repository.GetEmployees(ctx, cnpj)
}
