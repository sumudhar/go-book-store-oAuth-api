package db

import (
	"github.com/sumudhar/go-book-store-oAuth-api/src/domain/access_token"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)


func NewRepository() DbRepository{
	return  &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{

}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken,*errors.RestErr){
	return nil,errors.NewInternalServerError("database connection not implemented yet.")
}
