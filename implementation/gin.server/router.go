package gin_server

import (
	"github.com/kulinsky/guess_word/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterHandler struct {
	ucHandler usecase.Handler
	logger    *zap.Logger
}

func NewRouter(i usecase.Handler, logger *zap.Logger) RouterHandler {
	return RouterHandler{
		ucHandler: i,
		logger:    logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	rH.mainRoutes(api)
	rH.gameRoutes(api)
}

func (rH RouterHandler) mainRoutes(api *gin.RouterGroup) {
	r := api.Group("")
	r.GET("", rH.getVersion)
}

func (rH RouterHandler) gameRoutes(api *gin.RouterGroup) {
	r := api.Group("games")
	r.POST("", rH.startNewGame)
	r.GET(":game_id", rH.getGameByID)
	r.GET(":game_id/stat", rH.getGameStatByID)
	r.POST(":game_id/guess", rH.gameGuess)
}

func (rH RouterHandler) getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "0.1.0"})
}
