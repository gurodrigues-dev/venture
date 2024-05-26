package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	return nil, nil
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

func (s *ResponsibleService) InterfaceToString(value interface{}) (*string, error) {
	switch v := value.(type) {
	case string:
		return &v, nil
	default:
		return nil, fmt.Errorf("value isn't string")
	}
}

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
