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

type ResponsibleService struct {
	responsiblerepository repository.ResponsibleRepository
}

func NewResponsibleService(repo repository.ResponsibleRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: repo,
	}
}

func (rs *ResponsibleService) CreateResponsible(ctx context.Context, responsbile *types.Responsible) error {
	responsbile.Password = utils.HashPassword(responsbile.Password)
	return rs.responsiblerepository.CreateResponsible(ctx, responsbile)
}

func (rs *ResponsibleService) ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error) {
	return rs.responsiblerepository.ReadResponsible(ctx, cpf)
}

func (rs *ResponsibleService) UpdateResponsible(ctx context.Context) error {
	return rs.responsiblerepository.UpdateResponsible(ctx)
}

func (rs *ResponsibleService) DeleteResponsible(ctx context.Context, cpf *string) error {
	return rs.responsiblerepository.DeleteResponsible(ctx, cpf)
}

func (rs *ResponsibleService) AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error) {
	return nil, nil
}

func (rs *ResponsibleService) ParserJwtResponsible(ctx *gin.Context) (interface{}, error) {

	cpf, found := ctx.Get("cpf")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cpf, nil

}

func (rs *ResponsibleService) CreateTokenJWTResponsible(ctx context.Context, responsible *types.Responsible) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cpf": responsible.CPF,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil
}
