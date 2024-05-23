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
	schoolrepository repository.SchoolRepository
}

func NewSchoolService(repo repository.SchoolRepository) *SchoolService {
	return &SchoolService{
		schoolrepository: repo,
	}
}

func (ss *SchoolService) CreateSchool(ctx context.Context, school *types.School) error {
	school.Password = utils.HashPassword(school.Password)
	return ss.schoolrepository.CreateSchool(ctx, school)
}

func (ss *SchoolService) ReadSchool(ctx context.Context, cnpj *string) (*types.School, error) {
	return ss.schoolrepository.ReadSchool(ctx, cnpj)
}

func (ss *SchoolService) ReadAllSchools(ctx context.Context) ([]types.School, error) {
	return ss.schoolrepository.ReadAllSchools(ctx)
}

func (ss *SchoolService) UpdateCreateSchool(ctx context.Context) error {
	return ss.schoolrepository.UpdateSchool(ctx)
}

func (ss *SchoolService) DeleteSchool(ctx context.Context, cnpj *string) error {
	return ss.schoolrepository.DeleteSchool(ctx, cnpj)
}

func (ss *SchoolService) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	school.Password = utils.HashPassword(school.Password)
	return ss.schoolrepository.AuthSchool(ctx, school)
}

func (ss *SchoolService) CreateInvite(ctx context.Context, invite *types.Invite) error {
	return ss.schoolrepository.CreateInvite(ctx, invite)
}

func (ss *SchoolService) GetEmployees(ctx context.Context, cnpj *string) ([]types.Driver, error) {
	return ss.schoolrepository.GetEmployees(ctx, cnpj)
}

func (ss *SchoolService) IsEmployee(ctx context.Context, cnh *string) error {
	return ss.schoolrepository.IsEmployee(ctx, cnh)
}

func (ss *SchoolService) DeleteEmployee(ctx context.Context, record_id *int) error {
	return ss.schoolrepository.DeleteEmployee(ctx, record_id)
}

func (ss *SchoolService) ParserJwtSchool(ctx *gin.Context) (interface{}, error) {

	cnpj, found := ctx.Get("cnpj")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnpj, nil

}

func (ss *SchoolService) CreateTokenJWTSchool(ctx context.Context, school *types.School) (string, error) {

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
