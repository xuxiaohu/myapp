package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/xuxiaohu/myapp/internal/config"
	"github.com/xuxiaohu/myapp/internal/logger"
	"github.com/xuxiaohu/myapp/pkg/version"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		zap.L().Error("failed to load config", zap.Error(err))
		os.Exit(1)
	}

	log, err := logger.New(cfg.Env)
	if err != nil {
		zap.L().Error("failed to init logger", zap.Error(err))
		os.Exit(1)
	}
	defer log.Sync() //nolint:errcheck

	zap.ReplaceGlobals(log)

	log.Info("starting", zap.String("version", version.Version), zap.String("env", cfg.Env))
}
