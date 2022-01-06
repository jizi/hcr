package main

import (
	"context"
	"fmt"
	"github.com/pete911/hcr/internal/flag"
	"github.com/pete911/hcr/internal/hcr"
	"github.com/pete911/hcr/internal/logger"
	"go.uber.org/zap/zapcore"
	"os"
)

var Version = "dev"

func main() {
	log, err := logger.NewZapLogger(zapcore.InfoLevel)
	if err != nil {
		fmt.Printf("new zap logger: %v", err)
		os.Exit(1)
	}

	config, err := flag.ParseFlags()
	if err != nil {
		return
	}
	if config.Version {
		fmt.Println(Version)
		os.Exit(0)
	}

	log.Info(config.String())
	releaser, err := hcr.NewReleaser(log, config)
	if err != nil {
		log.Fatal(fmt.Sprintf("new releaser: %v", err))
	}
	if err := releaser.Release(context.TODO()); err != nil {
		log.Fatal(fmt.Sprintf("release: %v", err))
	}
}
