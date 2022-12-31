package urlservice

import (
	"net/http"
	"time"

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

}

func (uc UrlService) routeTo(ctx *gin.Context) {

	var url models.URL

	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller, exists := ctx.Get("UrlController")
	if !exists {
		ctx.Status(500)
		return
	}
	controller, ok := controller.(*mongourl.UrlController)
	if !ok {
		ctx.Status(500)
		return
	}

	ctx.JSON(200, "JSON")
}

func addDays(days int) int64 {
	return time.Now().AddDate(0, 0, days).Unix()
}
