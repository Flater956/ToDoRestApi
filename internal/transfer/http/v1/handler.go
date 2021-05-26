package v1

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sing_up", h.singUp)
		auth.POST("/sing_in", h.singIn)
	}

	lists := api.Group("/lists")
	{
		lists.GET("/", h.getLists)
		lists.POST("/", h.createList)
		lists.GET("/:id", h.getListById)
		lists.PUT("/:id", h.updateList)
		lists.DELETE("/:id", h.deleteList)

		items := lists.Group("/:id/items")
		{
			items.GET("/", h.getItems)
			items.POST("/", h.createItem)
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}
}
