package gin_server

import (
	"github.com/gin-gonic/gin"
	"github.com/kulinsky/guess_word/domain"
	"github.com/kulinsky/guess_word/implementation/formatter"
	"net/http"
)

type guessPostRequest struct {
	Letter string `json:"letter"`
}

func (rH RouterHandler) getGameStatByID(c *gin.Context) {
	gameID, err := domain.GameIDFromString(c.Param("game_id"))

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	gameStat, err := rH.ucHandler.GetGameStat(c, gameID)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"game": formatter.NewGameStatResponseFromDomain(gameStat)})
}

func (rH RouterHandler) getGameByID(c *gin.Context) {
	gameID, err := domain.GameIDFromString(c.Param("game_id"))

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	game, err := rH.ucHandler.GetGameByID(c, gameID)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"game": formatter.NewGameResponseFromDomain(game)})
}

func (rH RouterHandler) startNewGame(c *gin.Context) {
	word, err := rH.ucHandler.GetRandomWord(c)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	gameID, err := rH.ucHandler.StartNewGame(c, 5, word)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"game_id": gameID})
}

func (rH RouterHandler) gameGuess(c *gin.Context) {
	gameID, err := domain.GameIDFromString(c.Param("game_id"))

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	game, err := rH.ucHandler.GetGameByID(c, gameID)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	req := &guessPostRequest{}

	if err := c.BindJSON(req); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if req.Letter == "" {
		_ = c.Error(domain.ErrNotLetter)
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrNotLetter.Error()})

		return
	}

	guess, err := rH.ucHandler.Guess(c, &game.ID, req.Letter)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"guessed": guess})
}
