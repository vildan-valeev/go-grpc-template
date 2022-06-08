package v1

import (
	"github.com/gin-gonic/gin"
	"go-bolvanka/internal/service"

	"go-bolvanka/pkg/logging"
)

func NewRouter(handler *gin.Engine, l *logging.Logger, c service.CategoryService, i service.ItemService) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	//handler.GET("/swagger/*any", swaggerHandler)
	//
	//// K8s probe
	//handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	//
	//// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/api/v1")
	{
		newCategoryRoutes(h, c, *l)
		newItemRoutes(h, i, *l)
	}
}
