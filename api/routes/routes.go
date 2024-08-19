package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/gitignore-service/api/services"
)

var (
	StartTime time.Time
	BootTime  time.Duration
)

func InitRoutes(routes *gin.Engine) {
	// Initialize services
	svc := services.NewServices()

	// CORS middleware
	config := cors.DefaultConfig()
	// config = cors.Config{
	// 	AllowOrigins:     []string{"https://foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }
	config.AllowAllOrigins = true
	routes.Use(cors.New(config))

	// Serve the frontend
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	routes.LoadHTMLGlob("templates/**/*")
	// serve static pages under static folder
	routes.GET("/static/*filepath", func(ctx *gin.Context) {
		ctx.File("templates/static/" + ctx.Param("filepath"))
	})
	routes.GET("/media/*filepath", func(ctx *gin.Context) {
		ctx.File("templates/media/" + ctx.Param("filepath"))
	})
	// Home Page
	routes.GET("/", func(ctx *gin.Context) {
		svc.TemplatesService().Index(ctx)
	})

	// Initialize the routes
	// API Routes
	api := routes.Group("/api")
	{
		gitignore := api.Group("/gi")
		{
			gitignore.GET("/", func(ctx *gin.Context) {
				svc.GitignoreService().List(ctx)
			})
			gitignore.GET("/get", func(ctx *gin.Context) {
				svc.GitignoreService().Get(ctx)
			})
		}

		api.GET("/health", func(ctx *gin.Context) {
			svc.HealthCheckService().Status(ctx)
		})
		api.GET("/health/full", func(ctx *gin.Context) {
			svc.HealthCheckService().HealthCheck(ctx, StartTime, BootTime)
		})
	}
}
