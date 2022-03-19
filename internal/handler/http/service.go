package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/TestardR/platformmanager/internal/domain"
	"github.com/gin-gonic/gin"
)

var (
	errQueryNamespace = errors.New("failed to query namespace")
)

// @Summary GetServices queries pods information from cluster
// @Description Return services information
// @Tags Service
// @Produce json
// @Success 200 {object} []domain.Service
// @Failure 500 {object} ResponseError
// @Router /services [get].
func (h *handler) GetServices(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	defer cancel()

	dpls, err := h.pm.GetDeployments(ctx, h.cfg.Namespace)
	if err != nil {
		err = fmt.Errorf("%w: %s", errQueryNamespace, err)
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, newResponseError(err))
	}

	services := make([]domain.Service, 0, len(dpls))
	for _, dpl := range dpls {
		services = append(services, domain.NewServiceFromDpl(dpl))
	}

	c.JSON(http.StatusOK, services)
}

// @Summary GetServicesPerApplicationGroup queries pods information per application group
// @Description Return services information per application group
// @Tags Service
// @Produce json
// @Success 200 {object} []domain.Service
// @Failure 500 {object} ResponseError
// @Router /services [get].
func (h *handler) GetServicesPerApplicationGroup(c *gin.Context) {
	appGroup := c.Param("applicationGroup")
	h.log.Info(fmt.Sprintf("successfully called endpoint with applicationGroup: %s", appGroup))

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	defer cancel()

	dpls, err := h.pm.GetDeploymentsPerLabel(ctx, h.cfg.Namespace, "applicationGroup", appGroup)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, newResponseError(err))
	}

	services := make([]domain.Service, 0, len(dpls))
	for _, dpl := range dpls {
		services = append(services, domain.NewServiceFromDpl(dpl))
	}

	c.JSON(http.StatusOK, services)
}
