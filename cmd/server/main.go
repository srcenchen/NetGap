package main

import (
	"NetGap/internal/sdk"
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	srv, err := sdk.NewServer(
		sdk.WithServerTunnelAddr(":6500"))
	if err != nil {
		panic(err)
	}

	srv.Run(ctx)
}
