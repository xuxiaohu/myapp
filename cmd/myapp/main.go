package main

import (
	"fmt"
	"os"

	"github.com/xuxiaohu/myapp/internal/config"
	"github.com/xuxiaohu/myapp/pkg/version"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("myapp %s (env: %s)\n", version.Version, cfg.Env)
}
