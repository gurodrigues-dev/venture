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

type SchoolService struct {
	schoolrepository repository.SchoolRepositoryInterface
}

func NewSchoolService(repo repository.SchoolRepositoryInterface) *SchoolService {
	return &SchoolService{schoolrepository: repo}
}

func (s *SchoolService) CreateSchool(ctx context.Context, school *types.School) error {
	school.Password = utils.HashPassword(school.Password)
	return s.schoolrepository.CreateSchool(ctx, school)
}

func (s *SchoolService) ReadSchool(ctx context.Context, cnpj *string) (*types.School, error) {
	return s.schoolrepository.ReadSchool(ctx, cnpj)
}

func (s *SchoolService) ReadAllSchools(ctx context.Context) ([]types.School, error) {
	return s.schoolrepository.ReadAllSchools(ctx)
}

func (s *SchoolService) UpdateSchool(ctx context.Context, school *types.School) error {
	return s.schoolrepository.UpdateSchool(ctx, school)
}

func (s *SchoolService) DeleteSchool(ctx context.Context, cnpj *string) error {
	return s.schoolrepository.DeleteSchool(ctx, cnpj)
}

func (s *SchoolService) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	school.Password = utils.HashPassword(school.Password)
	return s.schoolrepository.AuthSchool(ctx, school)
}

func (s *SchoolService) CreateInvite(ctx context.Context, invite *types.Invite) error {
	return s.schoolrepository.CreateInvite(ctx, invite)
}

func (s *SchoolService) GetEmployees(ctx context.Context, cnpj *string) ([]types.Driver, error) {
	return s.schoolrepository.GetEmployees(ctx, cnpj)
}

func (s *SchoolService) IsEmployee(ctx context.Context, cnh *string) error {
	return s.schoolrepository.IsEmployee(ctx, cnh)
}

func (s *SchoolService) DeleteEmployee(ctx context.Context, record_id *int) error {
	return s.schoolrepository.DeleteEmployee(ctx, record_id)
}

func (s *SchoolService) ParserJwtSchool(ctx *gin.Context) (interface{}, error) {

	cnpj, found := ctx.Get("cnpj")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnpj, nil

}

func (s *SchoolService) CreateTokenJWTSchool(ctx context.Context, school *types.School) (string, error) {

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
