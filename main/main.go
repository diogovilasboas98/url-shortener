package main

import (
	b64 "encoding/base64"
	"link-shortner/database"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func Homepage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "homepage.tmpl", gin.H{})
}

func RedirectToURL(ctx *gin.Context) {
	id_b64 := ctx.Param("id")
	var link database.Link
	bDec, err := b64.URLEncoding.DecodeString(id_b64)

	if err != nil {
		log.Default().Print(err)
	}
	uDec, err := uuid.FromBytes(bDec)
	if err != nil {
		log.Default().Print(err)
	}

	database.DB.Where("ID = ?", uDec).First(&link)
	//log.Default().Print(link.URL)
	ctx.HTML(http.StatusOK, "redirect.tmpl", gin.H{"url": link.URL})
}

func CreateLink(ctx *gin.Context) {
	body := database.Link{}

	if err := ctx.ShouldBind(&body); err != nil {
		log.Default().Print(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	database.DB.Create(&body)
	bytes, err := body.ID.MarshalBinary()
	if err != nil {
		log.Default().Print(err)
	}
	sEnc := b64.URLEncoding.EncodeToString(bytes)
	//log.Default().Print(ctx.Request.Host + "/" + sEnc)

	ctx.HTML(http.StatusOK, "showURL.tmpl", gin.H{"url": ctx.Request.Host + "/" + url.QueryEscape(sEnc)})

}

func main() {
	godotenv.Load(".env")
	database.Connect()

	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.GET("/", Homepage)
	r.POST("/", CreateLink)
	r.GET("/:id", RedirectToURL)
	r.Run(":" + os.Getenv("PORT"))
}
