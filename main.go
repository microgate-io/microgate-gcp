package main

import (
	"context"

	"github.com/emicklei/microgate-io/microgate-gcp/v1/queue"
	"github.com/microgate-io/microgate"
	apilog "github.com/microgate-io/microgate-lib-go/v1/log"
	"github.com/microgate-io/microgate/v1/config"
	mconfig "github.com/microgate-io/microgate/v1/config"
	mlog "github.com/microgate-io/microgate/v1/log"
)

func main() {
	mlog.Init()

	gateConfig := config.Load("config.yaml")
	apilog.GlobalDebug, _ = gateConfig.FindBool("global_debug")

	qService, err := queue.NewQueueingServiceImpl(gateConfig)
	if err != nil {
		mlog.Fatalw(context.Background(), "cannot create queueing service", "err", err)
	}

	// these are the gRPC services provided to the backend
	provider := microgate.ServiceProvider{
		Log:      mlog.NewLogService(),
		Config:   mconfig.NewConfigServiceImpl(),
		Queueing: qService,
	}

	microgate.Start(gateConfig, provider)
}
