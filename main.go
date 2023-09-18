package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Van Halen", Artist: "Van Halen", Price: 16.99},
	{ID: "2", Title: "Scream Aim Fire", Artist: "Bullet For My Valentine", Price: 16.99},
	{ID: "3", Title: "Avenged Sevenfold", Artist: "Avenged Sevenfold", Price: 17.99},
}

// @title           Album API
// @version         1.0
// @description     This is Devon VanNote's first Go API!
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  dvon5150@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {	
    router := gin.Default()

    v1 := router.Group("/api/v1")
	{
		album := v1.Group("/album")
		{
			album.GET("", getAlbums)
			album.GET(":id", getAlbumByID)
			album.POST("", postAlbums)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Data not in correct format.")
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}