package config_test

import (
	"os"
	"testing"

	"github.com/xuxiaohu/myapp/internal/config"
)

func TestLoad_defaults(t *testing.T) {
	os.Unsetenv("APP_ENV")
	os.Unsetenv("APP_PORT")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Env != "development" {
		t.Errorf("expected env=development, got %q", cfg.Env)
	}
	if cfg.Port != "8080" {
		t.Errorf("expected port=8080, got %q", cfg.Port)
	}
}

func TestLoad_fromEnv(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("APP_PORT", "9090")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Env != "production" {
		t.Errorf("expected env=production, got %q", cfg.Env)
	}
	if cfg.Port != "9090" {
		t.Errorf("expected port=9090, got %q", cfg.Port)
	}
}
