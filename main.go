package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func authMiddleware(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	if len(auth) == 0 || !check(auth) {
		c.AbortWithStatusJSON(401, gin.H{"error": "Authorization failed"})
		return
	}

	c.Next()
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("using mongo on: " + os.Getenv("MONGO_URL"))
	log.Println("secret set to: " + os.Getenv("SECRET"))
}

func main() {

	r := gin.Default()

	r.Use(authMiddleware)

	r.GET("/:db/tagsCount", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query().Get("query")
		db := ctx.Param("db")

		count := getTagsCount(db, query)

		ctx.JSON(200, gin.H{
			"count": count,
		})
	})

	r.GET("/:db/tags", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query().Get("query")
		limit, _ := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
		skip, _ := strconv.Atoi(ctx.Request.URL.Query().Get("skip"))
		db := ctx.Param("db")

		tags := getTags(db, query, int64(limit), int64(skip))

		ctx.JSON(200, tags)
	})

	r.GET("/:db/gifs", func(ctx *gin.Context) {
		tag := ctx.Request.URL.Query().Get("tag")
		limit, _ := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
		skip, _ := strconv.Atoi(ctx.Request.URL.Query().Get("skip"))
		db := ctx.Param("db")

		var gifs []Gif
		switch tag {
		case "RANDOM":
			gifs = getRandomGifs(db, int64(limit))
			break
		case "ALL":
			gifs = getGifsByTag(db, "", int64(limit), int64(skip))
			break
		default:
			gifs = getGifsByTag(db, tag, int64(limit), int64(skip))
			break
		}

		ctx.JSON(200, gifs)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
