package main

import (
	"Magang/controller"
	"Magang/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	articleService service.ArticleService =  service.New()
	articleController controller.ArticleController = controller.New(articleService)
)

func main()  {
	// inisialisasi
	server := gin.Default()

	server.LoadHTMLGlob("template/*.html")
	
	// route dan isinya
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"massage" : "berhasil jalan!!!",
		})
	})

	// Get
	apiRoutes :=server.Group("/api")
	{
		apiRoutes.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, articleController.FindAll())
		})
	
		apiRoutes.POST("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, articleController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/posts", articleController.ShowAll)
	}

	// create
	server.POST("/create", func(ctx *gin.Context) {
		articleController.SaveToDB(ctx)
		ctx.Redirect(http.StatusFound, "/form")
	})

	// route untuk from add
	server.GET("/form", articleController.ShowForm)

	// route menampilkan data
	server.GET("/read", articleController.FindAllFromDB)

	// route untuk delete data
	server.POST("/delete", func(ctx *gin.Context) {
		articleController.Delete(ctx)
		ctx.Redirect(http.StatusOK, "/read")
	})

	// route untuk update data
	server.POST("/update", func(ctx *gin.Context) {
		articleController.Update(ctx)
		ctx.Redirect(http.StatusOK, "/read")
	})

	// run servernya
	server.Run(":8080")
}