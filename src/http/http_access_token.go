package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumudhar/go-book-store-oAuth-api/src/domain/access_token"
)



type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accesstokenHandler struct{
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accesstokenHandler {
              service: service,
	}
}

func (h *accesstokenHandler) GetById(c *gin.Context)  {
	accessToken, err:= h.service.GetById(c.Param("access_token_id"))
	if err != nil{
		c.JSON(err.Status,err) 
		return
	}
	c.JSON(http.StatusOK,accessToken)

}