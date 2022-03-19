package http

import (
	"github.com/TestardR/platformmanager/config"
	"github.com/TestardR/platformmanager/pkg/k8sclient"
	"github.com/TestardR/platformmanager/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	healthRoute                  = "/health"
	serviceRoute                 = "/services"
	serviceApplicationGroupRoute = "/services/:applicationGroup"
)

// Handler is the base struct for dependency injection.
type handler struct {
	cfg config.Conf
	log logger.Logger
	pm  k8sclient.Managerer
}

// @title Platform Manager Rest Server
// @version 1.0
// @description This a k8s platform manager cluster

// @contact.name Romain Testard
// @contact.email romain.rtestard@gmail.com

// @host localhost:3000
func NewServer(env string, log logger.Logger, pm k8sclient.Managerer) *gin.Engine {
	h := handler{
		log: log,
		pm:  pm,
	}

	gin.SetMode(env)

	router := gin.New()
	router.Use(gin.Recovery())

	// useful for monitoring our service and CI/CD tools
	router.GET(healthRoute, h.Health)

	router.GET(serviceRoute, h.GetServices)
	router.GET(serviceApplicationGroupRoute, h.GetServicesPerApplicationGroup)

	return router
}
