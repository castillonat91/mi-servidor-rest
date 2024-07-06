package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID      string `json:"id"`
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Año     int
}

var albums = []album{
	{ID: "1", Titulo: "Familia", Artista: "Juan Lopez", Año: 2020},
	{ID: "2", Titulo: "Adios", Artista: "Juan Lopez", Año: 2020},
	{ID: "3", Titulo: "Lluvia", Artista: "Juan Lopez", Año: 2020},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbun album

	if error := c.BindJSON(&newAlbun); error != nil {
		return
	}

	albums = append(albums, newAlbun)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return

		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"mensage": "album no encontrado"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"mensaje": "álbum eliminado"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje": "álbum no encontrado"})
}

func putAlbumById(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.JSON(http.StatusOK, gin.H{"mensaje": "álbum actualizado"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"mensaje": "álbum no encontrado"})
}

// rutas//
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.PUT("/albums/:id", putAlbumById)

	router.Run("localhost:8080")
}
