package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stocktacking_backend/internal/service/fixtures"
)

type FixturesHandler struct {
	fixtureService *fixtures.FixtureService
}

func NewFixturesHandler(
	fixtureService *fixtures.FixtureService,
) *FixturesHandler {
	return &FixturesHandler{
		fixtureService: fixtureService,
	}
}

func (h FixturesHandler) RunFixtures(c *gin.Context) {
	EnableCors(&c.Writer)
	if err := h.fixtureService.RunFixtures(c); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"succeed": true})
}

func EnableCors(w *gin.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
