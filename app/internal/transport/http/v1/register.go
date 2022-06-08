package v1

import (
	"github.com/gin-gonic/gin"
	"go-bolvanka/internal/domain"
	"go-bolvanka/internal/service"
	"go-bolvanka/pkg/logging"
	"net/http"
)

type categoryRoutes struct {
	c service.CategoryService
	l logging.Logger
}

type itemRoutes struct {
	c service.ItemService
	l logging.Logger
}

// собираем роуты в в группу хендлеров
func newCategoryRoutes(handler *gin.RouterGroup, c service.CategoryService, l logging.Logger) {
	r := &categoryRoutes{c, l}

	h := handler.Group("/categories")
	{
		h.GET("/all", r.GetCategoriesRouter)
		//h.POST("/create", r.CreateCategoryRouter)
		//h.DELETE("/delete", r.DeleteCategoryRouter)
	}
}

// собираем роуты в в группу хендлеров
func newItemRoutes(handler *gin.RouterGroup, c service.ItemService, l logging.Logger) {
	r := &itemRoutes{c, l}

	h := handler.Group("/items")
	{
		h.GET("/all", r.GetItemsRouter)
		//h.POST("/create", r.CreateItemRouter)
		//h.DELETE("/delete", r.DeleteItemRouter)
	}
}

// структуры ответов
type categoryResponse struct {
	Categories []domain.Category `json:"categories"`
}
type itemResponse struct {
	Items []domain.Item `json:"items"`
}

// получаем все записи категорий
func (r *categoryRoutes) GetCategoriesRouter(c *gin.Context) {
	categories, err := r.c.AllCategories(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - GetCategoriesRouter")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, categoryResponse{categories})
}

// получаем все записи
func (r *itemRoutes) GetItemsRouter(c *gin.Context) {
	items, err := r.c.AllItems(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, itemResponse{items})
}
