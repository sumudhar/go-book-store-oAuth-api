package access_token

import (
	"fmt"
	"strings"

	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)


type Repository interface{
	GetById(string)(*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	fmt.Println(accessTokenId)
    if len(accessTokenId)==0 {
		return nil, errors.NewBadRequestError("access token id is invalid")
	}
	accessToken, err:= s.repository.GetById(accessTokenId)
	if err != nil{
		return nil, err
	}
   return accessToken,nil
}
