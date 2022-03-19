package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthResp describes the response for Health Check.
type HealthResp struct {
	Status bool `json:"status"`
} // @name Health Response

// Health method http GET
// @Summary Health check
// @Description Healthcheck endpoint, to ensure that the service is running.
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthResp
// @Failure 500 {object} ResponseError
// @Router /health [get].
func (h *handler) Health(c *gin.Context) {
	// TODO: include health check for k8s cluster
	c.JSON(http.StatusOK, ResponseSuccess{Data: &HealthResp{Status: true}})
}
