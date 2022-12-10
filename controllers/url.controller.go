package controllers

import (
	"net/http"

	"example.com/urlibre/models"
	"example.com/urlibre/services"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	UrlService services.UrlService
}

func New(urlservice services.UrlService) UrlController {
	return UrlController{
		UrlService: urlservice,
	}
}

func (uc *UrlController) CreateUrl(ctx *gin.Context) {
	var url models.Url

	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UrlService.CreateUrl(&url)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UrlController) GetUrl(ctx *gin.Context) {
	urlid := ctx.Param("id")
	url, err := uc.UrlService.GetUser(&urlid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"URL": url})
}

func (uc *UrlController) RegisterUrlRoutes(rg *gin.RouterGroup) {
	ur := rg.Group("url")
	ur.POST("/create", uc.CreateUrl)
	ur.GET("/get/:id", uc.GetUrl)
}
