package urlservice

import (
	"net/http"

	"example.com/urlibre/models"
	mongourl "example.com/urlibre/mongo_url"
	"github.com/gin-gonic/gin"
)

type UrlService struct {
	urlController mongourl.UrlController
}

func (uc UrlService) New(us mongourl.UrlController) UrlService {
	return UrlService{us}
}

func (uc UrlService) AttachUrlServices(r *gin.Engine) *gin.Engine {
	r.GET("/:url_id", uc.routeTo)
	r.POST("/new", uc.newUrl)

	return r
}

func (uc UrlService) newUrl(ctx *gin.Context) {
	var url models.NewURL

	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newurl, err := uc.urlController.InsertUrl(url.IntoURL())
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.JSON(200, newurl)
}

func (uc UrlService) routeTo(ctx *gin.Context) {
	urlID := ctx.Param("url_id")

	url, err := uc.urlController.GetUrl(urlID)
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Redirect(302, url.Url)
}
