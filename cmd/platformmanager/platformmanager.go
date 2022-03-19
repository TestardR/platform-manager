package main

import (
	"fmt"

	"github.com/TestardR/platformmanager/config"
	"github.com/TestardR/platformmanager/internal/handler/http"
	"github.com/TestardR/platformmanager/pkg/k8sclient"
	"github.com/TestardR/platformmanager/pkg/logger"
)

const appName = "platformmanager"

func main() {
	log := logger.NewLogger(appName)

	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	platformManager, err := k8sclient.New(c.KubeConfigPath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to instantiate kubernetes client: %w", err))
	}

	s := http.NewServer(c.Env, log, platformManager)
	if err := s.Run(":" + c.Port); err != nil {
		log.Fatal(err)
	}
}
