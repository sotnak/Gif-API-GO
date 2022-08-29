package main

import (
	"strconv"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")

		if len(auth) == 0 || !check(auth) {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Authorization failed"})
			return
		}

		ctx.Next()
	}
}

func init() {
	initEnv()
}

func main() {

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	r := gin.Default()

	r.Use(authMiddleware())

	r.GET("/:db/tagsCount", cache.CacheByRequestURI(memoryStore, Env.TagsCacheTime),
		func(ctx *gin.Context) {
			query := ctx.Request.URL.Query().Get("query")
			db := ctx.Param("db")

			count := getTagsCount(db, query)

			ctx.JSON(200, gin.H{
				"count": count,
			})
		})

	r.GET("/:db/tags", cache.CacheByRequestURI(memoryStore, Env.TagsCacheTime),
		func(ctx *gin.Context) {
			query := ctx.Request.URL.Query().Get("query")
			limit, _ := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
			skip, _ := strconv.Atoi(ctx.Request.URL.Query().Get("skip"))
			db := ctx.Param("db")

			tags := getTags(db, query, int64(limit), int64(skip))

			ctx.JSON(200, tags)
		})

	r.GET("/:db/gifs", cache.CacheByRequestURI(memoryStore, Env.GifsCacheTime),
		func(ctx *gin.Context) {
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
