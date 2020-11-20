package main

import (
	"github.com/jordan-lumley/a1pos/cmd/monitor"
	"github.com/jordan-lumley/a1pos/cmd/periphies"
	"github.com/jordan-lumley/a1pos/internal/config"
	"github.com/jordan-lumley/a1pos/internal/logger"
	"github.com/jordan-lumley/a1pos/internal/service"
)

func main() {
	config.Config()

	logger.Instance().Info("main() initialize")

	// monitor and periphies are totally different, therefore should NOT
	// be a shared interface/struct to use. If I need them to be the same then
	// make then under the same package
	go monitor.Execute()

	go periphies.Execute()

	service.Run()
}
