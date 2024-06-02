package service

import (
	"context"
	"fmt"
	"gin/config"
	"gin/internal/repository"
	"gin/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ResponsibleService struct {
	responsiblerepository repository.ResponsibleRepositoryInterface
}

func NewResponsibleService(repo repository.ResponsibleRepositoryInterface) *ResponsibleService {
	return &ResponsibleService{responsiblerepository: repo}
}

func (s *ResponsibleService) CreateResponsible(ctx context.Context, responsible *types.Responsible) error {
	return s.responsiblerepository.CreateResponsible(ctx, responsible)
}

func (s *ResponsibleService) ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error) {
	return s.responsiblerepository.ReadResponsible(ctx, cpf)
}

func (s *ResponsibleService) UpdateResponsible(ctx context.Context) error {
	return s.responsiblerepository.UpdateResponsible(ctx)
}

func (s *ResponsibleService) DeleteResponsible(ctx context.Context, cpf *string) error {
	return s.responsiblerepository.DeleteResponsible(ctx, cpf)
}

func (s *ResponsibleService) AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error) {
	return s.responsiblerepository.AuthResponsible(ctx, responsible)
}

func (s *ResponsibleService) CreateChild(ctx context.Context, child *types.Child) error {
	return s.responsiblerepository.CreateChild(ctx, child)
}

func (s *ResponsibleService) ReadChildren(ctx context.Context, cpf *string) ([]types.Child, error) {
	return s.responsiblerepository.ReadChildren(ctx, cpf)
}

func (s *ResponsibleService) UpdateChild(ctx context.Context, child *types.Child) error {
	return s.responsiblerepository.UpdateChild(ctx, child)
}

func (s *ResponsibleService) DeleteChild(ctx context.Context, rg *string) error {
	return s.responsiblerepository.DeleteChild(ctx, rg)
}

func (s *ResponsibleService) IsSponsor(ctx context.Context, rg *string) bool {
	return s.responsiblerepository.IsSponsor(ctx, rg)
}

func (s *ResponsibleService) ParserJwtResponsible(ctx *gin.Context) (interface{}, error) {

	cpf, found := ctx.Get("cpf")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cpf, nil

}

func (s *ResponsibleService) CreateTokenJWTResponsible(ctx context.Context, responsible *types.Responsible) (string, error) {

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
