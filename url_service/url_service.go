package urlservice

import (
	"net/http"
	"net/url"

	"example.com/urlibre/models"
	mongourl "example.com/urlibre/mongo_url"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	var urlS models.NewURL

	if err := ctx.ShouldBindJSON(&urlS); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := url.ParseRequestURI(urlS.Url)
	if err != nil {
		ctx.JSON(400, "The provided \"URL\" could not be parsed as a valid URL")
		return
	}
	// Check if the url contains any banned words
	passed, err := models.FilterByList(urlS.Url)
	if err != nil {
		ctx.Status(400)
		return
	}
	// If the url didn't pass then return a 400 status code response
	if !passed {
		ctx.JSON(400, "The provided url contains banned words!")
		return
	}

	newurl, err := uc.urlController.InsertUrl(urlS.IntoURL())
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
		if err == mongo.ErrNoDocuments {
			ctx.String(404, "404 NOT FOUND!")
			return
		}
		ctx.Status(500)
		return
	}

	ctx.Redirect(302, url.Url)
}
